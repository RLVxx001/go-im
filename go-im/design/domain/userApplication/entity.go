package userApplication

import (
	"gorm.io/gorm"
	"time"
)

type UserApplication struct {
	gorm.Model
	UserOwner   uint      //所属用户或群id
	Class       uint      //0表示用户申请用户 1表示用户申请群 2表示群邀请用户
	Target      uint      //对方用户或群id
	InviteUser  uint      //当该请求为群邀请时邀请人id
	Remarks     string    `gorm:"type:varchar(500)"` //所属备注
	Text        string    `gorm:"type:varchar(500)"` //所属内容
	FailureTime time.Time //失效时间
	Stats       uint      //0代表申请 1表示拒绝 2表示接受
}

func NewUserApplication(userOwner, class, target uint, remarks, text string) *UserApplication {
	return &UserApplication{
		UserOwner: userOwner,
		Class:     class,
		Target:    target,
		Remarks:   remarks,
		Text:      text,
	}
}
