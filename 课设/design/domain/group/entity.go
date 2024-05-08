package group

import (
	"design/domain/groupMessage"
	"design/domain/groupUser"
	"gorm.io/gorm"
)

// Group 群表
type Group struct {
	gorm.Model
	GroupId       string
	GroupName     string
	GroupInform   string
	GroupUsers    []groupUser.GroupUser `gorm:"foreignKey:GroupId"`
	GroupMessages []groupMessage.GroupMessage
}

func NewGroup(groupId, groupName, groupInform string) *Group {
	return &Group{
		GroupId:     groupId,
		GroupName:   groupName,
		GroupInform: groupInform,
	}
}
