package userMessage

import (
	"gorm.io/gorm"
)

// UserMessage 用户消息
type UserMessage struct {
	gorm.Model
	UsertoUserId uint   //(用户-用户)表id
	Message      string `gorm:"type:varchar(500)"` //消息
	UserId       uint   //所属用户id
}

func NewUserMessage(usertoUserId, userId uint, message string) *UserMessage {
	return &UserMessage{
		UsertoUserId: usertoUserId,
		UserId:       userId,
		Message:      message,
	}
}
