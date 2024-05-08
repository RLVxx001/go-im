package groupMessage

import (
	"gorm.io/gorm"
)

// GroupMessage 群消息表
type GroupMessage struct {
	gorm.Model
	MessageOwner    uint   //消息所属用户
	MessageSender   uint   //消息发送用户
	MessageReceiver uint   //消息接收群
	Message         string `gorm:"type:varchar(500)"` //消息
}

func NewGroupUser(messageOwner, messageSender, messageReceiver uint, message string) *GroupMessage {
	return &GroupMessage{
		MessageOwner:    messageOwner,
		MessageSender:   messageSender,
		MessageReceiver: messageReceiver,
		Message:         message,
	}
}
