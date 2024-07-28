package usertoUser

import (
	"design/api/user"
	"design/domain/usertoUser"
	"time"
)

// 创建用户请求结构体
type UserRequest struct {
	UserOwner    uint          `json:"userOwner"`  //所属用户
	UserTarget   uint          `json:"userTarget"` //接受用户id
	Remarks      string        `json:"remarks"`    //备注
	Remarks1     string        `json:"remarks1"`   //备注1
	IsDeleted    bool          `json:"isDeleted"`  //是否被删除
	Shielded     bool          `json:"shielded"`   //是否被拉黑
	Message      string        `json:"message"`    //消息
	UserMessages []UserMessage `json:"userMessages"`
}

type UserMessage struct {
	Message      string    `json:"message"`      //消息
	UsertoUserId uint      `json:"usertoUserId"` //所属用户-用户id
	Key          uint      `json:"key"`          //消息标识
	User         uint      `json:"user"`         //消息消费者id
	UserOwner    uint      `json:"userOwner"`    //消息发送者id
	IsDeleted    bool      `json:"isDeleted"`    //是否被删除
	IsRead       bool      `json:"isRead"`       //是否被读
	CreatedAt    time.Time `json:"createdAt"`
}

// 创建用户响应
type UserResponse struct {
	ID           uint               `json:"id"`         //频道id
	UserOwner    uint               `json:"userOwner"`  //所属用户
	UserTarget   uint               `json:"userTarget"` //接受用户id
	Remarks      string             `json:"remarks"`    //备注
	IsDeleted    bool               `json:"isDeleted"`  //是否被删除
	Shielded     bool               `json:"shielded"`   //是否被拉黑
	ToUser       user.LoginResponse `json:"ToUser"`     //对方信息
	UserMessages []UserMessage      `json:"userMessages"`
}

// 类型转化
func ToUserResponse(u *usertoUser.UsertoUser) UserResponse {
	return UserResponse{
		ID:           u.ID,
		UserOwner:    u.UserOwner,
		UserTarget:   u.UserTarget,
		Remarks:      u.Remarks,
		IsDeleted:    u.IsDeleted,
		Shielded:     u.Shielded,
		UserMessages: ToUserMessage(u.UserMessages),
	}
}

func ToUsertoUser(u UserRequest) *usertoUser.UsertoUser {
	return &usertoUser.UsertoUser{
		UserOwner:    u.UserOwner,
		UserTarget:   u.UserTarget,
		Remarks:      u.Remarks,
		IsDeleted:    u.IsDeleted,
		Shielded:     u.Shielded,
		UserMessages: ToMessage(u.UserMessages),
	}
}

// 数组转化
func ToUserMessage(us []usertoUser.UserMessage) []UserMessage {
	users := make([]UserMessage, len(us))
	for i, j := range us {
		users[i] = UserMessage{
			Message:      j.Message,
			UsertoUserId: j.UsertoUserId,
			Key:          j.Key,
			CreatedAt:    j.CreatedAt,
			UserOwner:    j.UserOwner,
			IsDeleted:    j.IsDeleted,
			IsRead:       j.IsRead,
		}

	}
	return users
}

func ToMessage(us []UserMessage) []usertoUser.UserMessage {
	users := make([]usertoUser.UserMessage, len(us))
	for i, j := range us {
		users[i] = usertoUser.UserMessage{
			Message:      j.Message,
			UsertoUserId: j.UsertoUserId,
			Key:          j.Key,
			IsRead:       j.IsRead,
		}

	}
	return users
}
