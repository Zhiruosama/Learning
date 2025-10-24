package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Type       string `json:"type" gorm:"size:16;index;not null"`
	RoomID     uint   `json:"room_id" gorm:"index"`
	SenderID   uint   `json:"sender_id" gorm:"index;not null"`
	ReceiverID uint   `json:"receiver_id" gorm:"index"`
	SenderName string `json:"sender_name" gorm:"size:50;not null"`
	Content    string `json:"content" gorm:"type:text;not null"`
}

func SaveGroup(roomID, senderID uint, senderName, content string) error {
	m := Message{
		Type:       "group",
		RoomID:     roomID,
		SenderID:   senderID,
		SenderName: senderName,
		Content:    content,
	}
	return DB.Create(&m).Error
}

func SavePrivate(senderID, receiverID uint, senderName, content string) error {
	m := Message{
		Type:       "private",
		SenderID:   senderID,
		ReceiverID: receiverID,
		SenderName: senderName,
		Content:    content,
	}
	return DB.Create(&m).Error
}

func ListGroupMessages(roomID uint, limit int) ([]Message, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	var ms []Message
	err := DB.Where("type=? AND room_id=?", "group", roomID).Order("id desc").Limit(limit).Find(&ms).Error

	//倒序取出 返回时升序
	for i, j := 0, len(ms)-1; i < j; i, j = i+1, j-1 {
		ms[i], ms[j] = ms[j], ms[i]
	}
	return ms, err
}

func ListPrivateMessages(selfID, peerID uint, limit int) ([]Message, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	var ms []Message
	err := DB.Where("type=? AND ((sender_id=? AND receiver_id=?) OR (sender_id=? AND receiver_id=?))", "private", selfID, peerID, peerID, selfID).
		Order("id desc").Limit(limit).Find(&ms).Error

	//倒序取出 返回时升序
	for i, j := 0, len(ms)-1; i < j; i, j = i+1, j-1 {
		ms[i], ms[j] = ms[j], ms[i]
	}
	return ms, err
}
