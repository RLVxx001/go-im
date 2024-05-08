package space

import (
	"design/domain/spaceTrends"
	"gorm.io/gorm"
)

// Space 空间表
type Space struct {
	gorm.Model
	UserId      uint                      //用户id
	EnterPermit bool                      //是否设置禁入
	SpaceTrends []spaceTrends.SpaceTrends `gorm:"foreignKey:SpaceId"` //空间动态[](不计入表)
}

func NewSpace(userId uint) *Space {
	return &Space{
		UserId:      userId,
		EnterPermit: false,
		SpaceTrends: make([]spaceTrends.SpaceTrends, 0),
	}
}
