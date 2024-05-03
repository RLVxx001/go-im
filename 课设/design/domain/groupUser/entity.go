package groupUser

import (
	"gorm.io/gorm"
)

// GroupUser 群用户表
type GroupUser struct {
	gorm.Model
	GroupId uint //群表id
	UserId  uint //用户id
	IsAdmin bool //是否管理
	IsGag   bool //是否被禁言
}

func NewGroupUser(groupId uint,userId uint) *GroupUser {
	return &GroupUser{
		GroupId: groupId,
		UserId: userId,
		IsAdmin: false,
		IsGag: false,
	}
}
