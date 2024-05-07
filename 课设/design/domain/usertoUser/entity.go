package usertoUser

import (
	"gorm.io/gorm"
)

// UsertoUser 用户-用户表
type UsertoUser struct {
	gorm.Model
	UserOwner    uint   //所属用户
	UserTarget   uint   //接受用户id
	Remarks      string `gorm:"type:varchar(500)"` //备注
	IsDeleted    bool   //是否被删除
	Shielded     bool   //是否被拉黑
	UserMassages []UserMessage
}

// UserMessage 用户消息
type UserMessage struct {
	gorm.Model
	Message      string `gorm:"type:varchar(500)"` //消息
	UsertoUserId uint   //所属用户-用户id
	Key          uint   //消息标识
}

func NewUsertoUser(userOwner, userTarget uint, remarks string) *UsertoUser {
	return &UsertoUser{
		UserOwner:  userOwner,
		UserTarget: userTarget,
		Remarks:    remarks,
		Shielded:   false,
		IsDeleted:  false,
	}
}

func NewUserMessage(utouid uint, message string) *UserMessage {
	return &UserMessage{
		Message:      message,
		UsertoUserId: utouid,
	}
}
