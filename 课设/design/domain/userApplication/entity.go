package userApplication

import (
	"gorm.io/gorm"
)

type UserApplication struct {
	gorm.Model
	UserOwner  uint   //所属用户
	UserTarget uint   //对方用户
	Remarks    string `gorm:"type:varchar(500)"` //所属备注
	Text       string `gorm:"type:varchar(500)"` //所属内容
	IsAccept   bool   //是否接受
	IsDown     bool   //是否拒绝
}

func NewUserApplication(userOwner, userTarget uint, remarks, text string) *UserApplication {
	return &UserApplication{
		UserOwner:  userOwner,
		UserTarget: userTarget,
		Remarks:    remarks,
		Text:       text,
		IsAccept:   false,
	}
}
