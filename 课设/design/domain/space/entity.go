package space

import (
	"design/domain/user"
	"gorm.io/gorm"
)

// Space 空间表
type Space struct {
	gorm.Model
	UserId      uint          //用户id
	EnterPermit bool          //是否设置禁入
	SpaceTrends []SpaceTrends `gorm:"foreignKey:SpaceId"` //空间动态[](不计入表)
}

func NewSpace(userId uint) *Space {
	return &Space{
		UserId:      userId,
		EnterPermit: false,
		SpaceTrends: make([]SpaceTrends, 0),
	}
}

type SpaceTrends struct {
	gorm.Model
	UserId   uint
	User     user.User `gorm:"-"`
	Detail   string
	Praise   uint
	Comments []Comment `gorm:"foreignKey:SpaceTrendsId"` //评论[](不计入表)
	SpaceId  uint      //空间表id
}

func NewSpaceTrends(spaceId uint) *SpaceTrends {
	return &SpaceTrends{
		SpaceId:  spaceId,
		Comments: make([]Comment, 0),
	}
}

// Comment 评论表
type Comment struct {
	gorm.Model
	UserId        uint      //评论用户id
	User          user.User `gorm:"-"`
	Praise        uint
	Content       string //内容
	SpaceTrendsId uint   //空间动态表id
	ToUserId      uint   //0
}

func NewComment(userId uint, content string, spaceTrendsId uint) *Comment {
	return &Comment{
		UserId:        userId,
		Content:       content,
		SpaceTrendsId: spaceTrendsId,
	}
}
