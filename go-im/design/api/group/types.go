package group

import (
	"design/domain/group"
	"design/domain/user"
	"time"
)

type GroupRequest struct {
	Id            uint           `json:"id"`            //群id
	GroupId       string         `json:"groupId"`       //群号
	GroupName     string         `json:"groupName"`     //群名称
	GroupInform   string         `json:"groupInform"`   //群公告
	GroupUsers    []GroupUser    `json:"groupUsers"`    //群用户
	GroupMessages []GroupMessage `json:"groupMessages"` //群消息
	Img           string         `json:"img"`           //群头像
	UpdatedAt     time.Time      `json:"updatedAt"`     //更新事件
}

// GroupMessage 群消息表
type GroupMessage struct {
	Id            uint      `json:"id"`            //群消息id
	MessageOwner  uint      `json:"messageOwner"`  //消息所属用户
	MessageSender uint      `json:"messageSender"` //消息发送用户
	SenderUser    GroupUser `json:"senderUser"`    //消息发送用户实体
	GroupId       uint      `json:"groupId"`       //消息接收群
	Message       string    `json:"message"`       //消息
	MessageKey    uint      `json:"messageKey"`    //消息key
	IsRead        bool      `json:"isRead"`        //是否已读
	Img           string    `json:"img"`           //图片
	IsDeleted     bool      `json:"isDeleted"`     //是否被删除
	UpdatedAt     time.Time `json:"updatedAt"`
}

// GroupUser 群用户表
type GroupUser struct {
	ID      uint   `json:"id"`      //群用户id
	GroupId uint   `json:"groupId"` //群表id
	UserId  uint   `json:"userId"`  //用户id
	User    User   `json:"user"`
	IsAdmin uint   `json:"isAdmin"` //是否管理
	IsGag   bool   `json:"isGag"`   //是否被禁言
	Text    string `json:"text"`    //群备注
}
type User struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	Img      string `json:"img"`
}

func ToGroup(req GroupRequest) group.Group {
	return group.Group{
		GroupId:     req.GroupId,
		GroupName:   req.GroupName,
		GroupInform: req.GroupInform,
		Img:         req.Img,
	}
}
func ToGroupRequest(req group.Group) GroupRequest {
	return GroupRequest{
		Id:            req.ID,
		GroupId:       req.GroupId,
		GroupName:     req.GroupName,
		GroupInform:   req.GroupInform,
		GroupUsers:    ToResponseGroupUsers(req.GroupUsers),
		GroupMessages: ToResponseGroupMessages(req.GroupMessages),
		UpdatedAt:     req.UpdatedAt,
		Img:           req.Img,
	}
}

func ToGroupUser(req GroupUser) group.GroupUser {
	return group.GroupUser{
		ID:      req.ID,
		GroupId: req.GroupId,
		UserId:  req.UserId,
		IsAdmin: req.IsAdmin,
		IsGag:   req.IsGag,
		Text:    req.Text,
	}
}

func ToResponseGroupUsers(req []group.GroupUser) []GroupUser {
	var gs []GroupUser
	for _, i := range req {
		gs = append(gs, ToResponseGroupUser(i))
	}
	return gs
}

func ToResponseGroupUser(req group.GroupUser) GroupUser {
	return GroupUser{
		ID:      req.ID,
		GroupId: req.GroupId,
		UserId:  req.UserId,
		User:    ToUser(req.User),
		IsAdmin: req.IsAdmin,
		IsGag:   req.IsGag,
		Text:    req.Text,
	}
}

func ToResponseGroupMessages(req []group.GroupMessage) []GroupMessage {
	var gs []GroupMessage
	for _, i := range req {
		gs = append(gs, ToResponseGroupMessage(i))
	}
	return gs
}
func ToResponseGroupMessage(req group.GroupMessage) GroupMessage {
	return GroupMessage{
		Id:            req.ID,
		MessageOwner:  req.MessageOwner,
		MessageSender: req.MessageSender,
		SenderUser:    ToResponseGroupUser(req.SenderUser),
		GroupId:       req.GroupId,
		Message:       req.Message,
		MessageKey:    req.MessageKey,
		IsRead:        req.IsRead,
		UpdatedAt:     req.UpdatedAt,
		IsDeleted:     req.IsDeleted,
	}
}
func ToUser(u user.User) User {
	return User{
		Username: u.Username,
		Account:  u.Account,
		Img:      u.Img,
	}
}
