package usertoUser

import "design/domain/usertoUser"

// 创建用户请求结构体
type UserRequest struct {
	UserOwner  uint   `json:"userOwner"`  //所属用户
	UserTarget uint   `json:"userTarget"` //接受用户id
	Remarks    string `json:"remarks"`    //备注
	IsDeleted  bool   `json:"isDeleted"`  //是否被删除
	Shielded   bool   `json:"shielded"`   //是否被拉黑
}

type UserMessage struct {
	Message      string `json:"message"`      //消息
	UsertoUserId uint   `json:"usertoUserId"` //所属用户-用户id
	Key          uint   `json:"key"`          //消息标识
}

// 创建用户响应
type UserResponse struct {
	UserOwner    uint          `json:"userOwner"`  //所属用户
	UserTarget   uint          `json:"userTarget"` //接受用户id
	Remarks      string        `json:"remarks"`    //备注
	IsDeleted    bool          `json:"isDeleted"`  //是否被删除
	Shielded     bool          `json:"shielded"`   //是否被拉黑
	UserMassages []UserMessage `json:"UserMassages"`
}

// 类型转化
func ToUserResponse(u *usertoUser.UsertoUser) UserResponse {
	return UserResponse{
		UserOwner:    u.UserOwner,
		UserTarget:   u.UserTarget,
		Remarks:      u.Remarks,
		IsDeleted:    u.IsDeleted,
		Shielded:     u.Shielded,
		UserMassages: ToUserMessage(u.UserMassages),
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
		}

	}
	return users
}
