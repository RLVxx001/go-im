package userApplication

import (
	"design/domain/userApplication"
	"time"
)

// 创建用户请求结构体
type CreateRequest struct {
	UserOwner  uint   `json:"userOwner"`  //所属用户或群id
	Class      uint   `json:"class"`      //0表示用户xx用户 1表示用户xx群 2表示群xx用户
	InviteUser uint   `json:"inviteUser"` //当该请求为群邀请时邀请人id
	Target     uint   `json:"target"`     //对方用户或群id
	Remarks    string `json:"remarks"`    //所属备注
	Text       string `json:"text"`       //所属内容
	Stats      uint   `json:"stats"`      //0代表申请 1表示拒绝 2表示接受
}

// 创建用户响应
type CreateResponse struct {
	UserOwner     uint          `json:"userOwner"`     //所属用户或群id
	Class         uint          `json:"class"`         //0表示用户申请用户 1表示用户申请群 2表示群邀请用户
	InviteUser    uint          `json:"inviteUser"`    //当该请求为群邀请时邀请人id
	Target        uint          `json:"target"`        //对方用户或群id
	Remarks       string        `json:"remarks"`       //所属备注
	Text          string        `json:"text"`          //所属内容
	Stats         uint          `json:"stats"`         //0代表申请 1表示拒绝 2表示接受
	UpdatedAt     time.Time     `json:"updatedAt"`     //更改时间
	FailureTime   time.Time     `json:"failureTime"`   //失效时间
	UserResponse  UserResponse  `json:"userResponse"`  //用户信息
	GroupResponse GroupResponse `json:"groupResponse"` //群信息
}

// 用户信息响应
type UserResponse struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	Img      string `json:"img"`
}

// 群信息响应
type GroupResponse struct {
	GroupId   string `json:"groupId"`   //群号
	GroupName string `json:"groupName"` //群名称
	Img       string `json:"img"`
}

func ToCreateResponse(u userApplication.UserApplication) CreateResponse {
	return CreateResponse{
		UserOwner:   u.UserOwner,
		Target:      u.Target,
		Class:       u.Class,
		InviteUser:  u.InviteUser,
		Remarks:     u.Remarks,
		Text:        u.Text,
		Stats:       u.Stats,
		UpdatedAt:   u.UpdatedAt,
		FailureTime: u.FailureTime,
	}
}
