package ws

import (
	"encoding/json"
	"log"
	"minimal_chat/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type clientMsg struct {
	Status int               `json:"status"`
	Data   map[string]string `json:"data"`
}

type serveMsg struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type wsClient struct {
	Conn     *websocket.Conn
	Uid      string
	Username string
	RoomId   string
	AvatarId string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// 房间码 对应其 房间内的客户端
var rooms = make(map[int][]wsClient)

var clientsByConn = make(map[*websocket.Conn]wsClient)

var clientsByUid = make(map[string][]wsClient)

// 状态码 用于功能分发
const (
	msgTypeOnline        = 1
	msgTypeMembers       = 2
	msgTypeSend          = 3
	msgTypeGetOnlineUser = 4
	msgTypePrivate       = 5
)

// GET /ws
func Start(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer func() {
		conn.Close()
		cleanupConn(conn)
	}()

	for {
		mt, payload, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		//文本心跳
		if mt == websocket.TextMessage && string(payload) == "heartbeat" {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(`{
			"status":0,
			"data":"heartbeat ok"
			}`))
			continue
		}

		var cm clientMsg
		if err := json.Unmarshal(payload, &cm); err != nil {
			log.Println("json error:", err)
			continue
		}

		// 消息分发
		switch cm.Status {
		case msgTypeOnline:
			handleOnline(conn, cm)
		case msgTypeSend:
			handleSend(cm)
		case msgTypeGetOnlineUser:
			handleGetOnline(conn, cm)
		case msgTypePrivate:
			handlePrivate(conn, cm)
		case msgTypeMembers:
			handleMembers(conn, cm)
		default:
		}
	}
}

// 处理连接
func handleOnline(conn *websocket.Conn, cm clientMsg) {
	uid := cm.Data["uid"]
	username := cm.Data["username"]
	roomId := cm.Data["room_id"]
	avatar := cm.Data["avatar_id"]

	ri, _ := strconv.Atoi(roomId)

	//踢掉同uid的旧连接
	var next []wsClient
	for _, c := range rooms[ri] {
		if c.Uid == uid {
			_ = c.Conn.WriteMessage(websocket.TextMessage, []byte(`"status":-1,"data":[]`))
			_ = c.Conn.Close()
			continue
		}
		next = append(next, c)
	}
	next = append(next, wsClient{
		Conn:     conn,
		Uid:      uid,
		Username: username,
		RoomId:   roomId,
		AvatarId: avatar,
	})
	rooms[ri] = next

	// 登记连接索引
	clientsByConn[conn] = wsClient{
		Conn:     conn,
		Uid:      uid,
		Username: username,
		RoomId:   roomId,
		AvatarId: avatar,
	}
	clientsByUid[uid] = append(clientsByUid[uid], clientsByConn[conn])

	//回执上线
	_ = conn.WriteJSON(serveMsg{
		Status: msgTypeOnline,
		Data: map[string]interface{}{
			"uid": uid, "username": username, "room_id": roomId, "avatar_id": avatar,
			"time": time.Now().UnixNano(),
		},
	})
}

// 处理信息发送
func handleSend(cm clientMsg) {
	uid := cm.Data["uid"]
	username := cm.Data["username"]
	roomId := cm.Data["room_id"]
	avatar := cm.Data["avatar_id"]
	content := cm.Data["content"]

	msg := serveMsg{
		Status: msgTypeSend,
		Data: map[string]interface{}{
			"uid": uid, "username": username, "room_id": roomId, "avatar_id": avatar,
			"content": content, "time": time.Now().UnixNano(),
		},
	}

	//持久化群聊消息
	uid64, _ := strconv.ParseUint(uid, 10, 64)

	ri, _ := strconv.Atoi(roomId)
	_ = models.SaveGroup(uint(ri), uint(uid64), username, content)
	for _, c := range rooms[ri] {
		_ = c.Conn.WriteJSON(msg)
	}
}

// 获取房间内用户信息
func handleGetOnline(conn *websocket.Conn, cm clientMsg) {
	roomId := cm.Data["room_id"]
	ri, _ := strconv.Atoi(roomId)
	list := rooms[ri]

	type onlineItem struct {
		Uid      string `json:"uid"`
		Username string `json:"username"`
		AvatarId string `json:"avatar_id"`
		RoomId   string `json:"room_id"`
	}
	var data []onlineItem
	for _, c := range list {
		data = append(data, onlineItem{
			Uid: c.Uid, Username: c.Username, AvatarId: c.AvatarId, RoomId: c.RoomId,
		})
	}

	_ = conn.WriteJSON(serveMsg{
		Status: msgTypeGetOnlineUser,
		Data:   map[string]interface{}{"count": len(data), "list": data},
	})
}

func handlePrivate(conn *websocket.Conn, cm clientMsg) {
	sender, ok := clientsByConn[conn]
	if !ok {
		return
	}

	toUid := cm.Data["to_uid"]
	content := cm.Data["content"]
	if toUid == "" || content == "" {
		_ = conn.WriteJSON(serveMsg{Status: -1, Data: map[string]interface{}{
			"error": "to_uid and content are required",
		}})
		return
	}

	// 解析并校验接收方 uid
	tid64, err := strconv.ParseUint(toUid, 10, 64)
	if err != nil || tid64 == 0 {
		_ = conn.WriteJSON(serveMsg{Status: -1, Data: map[string]interface{}{"error": "invalid to_uid"}})
		return
	}
	sid64, _ := strconv.ParseUint(sender.Uid, 10, 64)

	// 持久化私信（确保 receiver_id 正确）
	if err := models.SavePrivate(uint(sid64), uint(tid64), sender.Username, content); err != nil {
		_ = conn.WriteJSON(serveMsg{Status: -1, Data: map[string]interface{}{"error": "save private failed"}})
		return
	}

	//构造回传信息（发件人和收件人都能收到）
	msg := serveMsg{
		Status: msgTypePrivate,
		Data: map[string]interface{}{
			"uid":        sender.Uid,
			"username":   sender.Username,
			"avatar_id":  sender.AvatarId,
			"to_uid":     toUid,
			"content":    content,
			"time":       time.Now().UnixNano(),
			"is_private": true,
		},
	}

	//发给对方
	if recvs, ok := clientsByUid[toUid]; ok {
		for _, rc := range recvs {
			_ = rc.Conn.WriteJSON(msg)
		}
	}

	_ = conn.WriteJSON(msg)
}

// 在断开连接的时候清除客户端信息
func cleanupConn(conn *websocket.Conn) {
	cl, ok := clientsByConn[conn]
	if !ok {
		return
	}
	delete(clientsByConn, conn)

	//从按用户索引中移除
	uid := cl.Uid
	var kept []wsClient
	for _, c := range clientsByUid[uid] {
		if c.Conn != conn {
			kept = append(kept, c)
		}
	}
	if len(kept) == 0 {
		delete(clientsByUid, uid)
	} else {
		clientsByUid[uid] = kept
	}

	//从房间列表中移除
	ri, _ := strconv.Atoi(cl.RoomId)
	var rkept []wsClient
	for _, c := range rooms[ri] {
		if c.Conn != conn {
			rkept = append(rkept, c)
		}
	}
	if len(rkept) == 0 {
		delete(rooms, ri)
	} else {
		rooms[ri] = rkept
	}
}

func handleMembers(conn *websocket.Conn, cm clientMsg) {
	roomId := cm.Data["room_id"]
	if roomId == "" {
		if me, ok := clientsByConn[conn]; ok {
			roomId = me.RoomId
		}
	}
	if roomId == "" {
		return
	}

	ri, err := strconv.Atoi(roomId)
	if err != nil {
		return
	}

	//去重聚合成员信息
	dedup := make(map[string]wsClient)
	for _, c := range rooms[ri] {
		dedup[c.Uid] = c
	}

	//构造列表
	var members []map[string]interface{}
	for _, m := range dedup {
		members = append(members, map[string]interface{}{
			"uid":       m.Uid,
			"username":  m.Username,
			"avatar_id": m.AvatarId,
			"room_id":   m.RoomId,
		})
	}

	//回传给请求方
	_ = conn.WriteJSON(serveMsg{
		Status: msgTypeMembers,
		Data: map[string]interface{}{
			"room_id": roomId,
			"count":   len(members),
			"members": members,
			"time":    time.Now().UnixNano(),
		},
	})
}
