package userApplication

import (
	"design/domain/userApplication"
	"time"
)

// 创建用户请求结构体
type CreateRequest struct {
	UserOwner  uint   `json:"userOwner"`  //所属用户
	UserTarget uint   `json:"userTarget"` //对方用户
	Remarks    string `json:"remarks"`    //所属备注
	Text       string `json:"text"`       //所属内容
	IsDown     bool   `json:"isDown"`     //是否拒绝
}

// 创建用户响应
type CreateResponse struct {
	UserOwner    uint         `json:"userOwner"`    //所属用户
	UserTarget   uint         `json:"userTarget"`   //对方用户
	Remarks      string       `json:"remarks"`      //所属备注
	Text         string       `json:"text"`         //所属内容
	IsAccept     bool         `json:"isAccept"`     //是否接受
	IsDown       bool         `json:"isDown"`       //是否拒绝
	UpdatedAt    time.Time    `json:"updatedAt"`    //更改时间
	UserResponse UserResponse `json:"userResponse"` //用户信息
}

// 用户信息响应
type UserResponse struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	Img      string `json:"img"`
}

func ToCreateResponse(u userApplication.UserApplication) CreateResponse {
	return CreateResponse{
		UserOwner:  u.UserOwner,
		UserTarget: u.UserTarget,
		Remarks:    u.Remarks,
		Text:       u.Text,
		IsAccept:   u.IsAccept,
		IsDown:     u.IsDown,
		UpdatedAt:  u.UpdatedAt,
	}
}
