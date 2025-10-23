package ws

import (
	"encoding/json"
	"log"
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

var rooms = make(map[int][]wsClient)

const (
	msgTypeOnline = 1
	msgTypeSend   = 3
)

// GET /ws
func Start(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer conn.Close()

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
		default:
		}
	}
}

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

	//回执上线
	_ = conn.WriteJSON(serveMsg{
		Status: msgTypeOnline,
		Data: map[string]interface{}{
			"uid": uid, "username": username, "room_id": roomId, "avatar_id": avatar,
			"time": time.Now().UnixNano(),
		},
	})
}

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

	ri, _ := strconv.Atoi(roomId)
	for _, c := range rooms[ri] {
		_ = c.Conn.WriteJSON(msg)
	}
}
