package commennt

import "gorm.io/gorm"

// Comment 评论表
type Comment struct {
	gorm.Model
	UserId        uint   //评论用户id
	Content       string //内容
	SpaceTrendsId uint   //空间动态表id
}

func NewSpace(userId uint, content string, spaceTrendsId uint) *Comment {
	return &Comment{
		UserId:        userId,
		Content:       content,
		SpaceTrendsId: spaceTrendsId,
	}
}
