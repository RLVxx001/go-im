package spaceTrends

import (
	"design/domain/commennt"
	"gorm.io/gorm"
)

// SpaceTrends 空间动态表
type SpaceTrends struct {
	gorm.Model
	Comments []commennt.Comment `gorm:"foreignKey:SpaceTrendsId"` //评论[](不计入表)
	SpaceId  uint               //空间表id
}

func NewSpaceTrends(spaceId uint) *SpaceTrends {
	return &SpaceTrends{
		SpaceId:  spaceId,
		Comments: make([]commennt.Comment, 0),
	}
}
