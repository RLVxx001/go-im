package usertoUser

import (
	"design/domain/userMessage"
	"gorm.io/gorm"
)

// UsertoUser 用户-用户表
type UsertoUser struct {
	gorm.Model
	UserOwner    uint   //(用户-用户)表id
	Remarks      string //消息
	Shielded     bool
	UserMassages []userMessage.UserMessage
}

func NewUsertoUser(userOwner uint, remarks string) *UsertoUser {
	return &UsertoUser{
		UserOwner: userOwner,
		Remarks:   remarks,
		Shielded:  false,
	}
}
