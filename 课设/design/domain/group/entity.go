package group

import (
	"design/domain/user"
	"gorm.io/gorm"
)

// Group 群表
type Group struct {
	gorm.Model
	GroupId       string         //群号
	GroupName     string         //群名称
	GroupInform   string         //群公告
	Img           string         //群头像
	GroupUsers    []GroupUser    `gorm:"foreignKey:GroupId"`
	GroupMessages []GroupMessage `gorm:"-"`
}

// GroupMessage 群消息表
type GroupMessage struct {
	gorm.Model
	MessageOwner  uint      //消息所属用户
	MessageSender uint      //消息发送用户
	SenderUser    user.User `gorm:"-"` //发送用户实体
	GroupId       uint      //消息接收群
	Message       string    `gorm:"type:varchar(500)"` //消息
	Img           string    `消息包含图片`
	MessageKey    uint      //消息key
	IsRead        bool      //是否已读
}

// GroupUser 群用户表
type GroupUser struct {
	ID      uint      `gorm:"primarykey"`
	GroupId uint      //群表id
	UserId  uint      //用户id
	IsAdmin uint      //是否管理
	IsGag   bool      //是否被禁言
	Text    string    //群备注
	User    user.User `gorm:"-"` //用户实体
}

func NewGroupUser(groupId uint, userId uint) *GroupUser {
	return &GroupUser{
		GroupId: groupId,
		UserId:  userId,
		IsAdmin: 0,
		IsGag:   false,
	}
}
func NewGroupMessage(messageOwner, messageSender, groupId uint, message string) *GroupMessage {
	return &GroupMessage{
		MessageOwner:  messageOwner,
		MessageSender: messageSender,
		GroupId:       groupId,
		Message:       message,
	}
}

func NewGroup(groupId, groupName, groupInform string) *Group {
	return &Group{
		GroupId:     groupId,
		GroupName:   groupName,
		GroupInform: groupInform,
	}
}
