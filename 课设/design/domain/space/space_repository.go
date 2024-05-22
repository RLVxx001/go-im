package space

import (
	"fmt"
	"gorm.io/gorm"
)

type SpaceRepostirory struct {
	db *gorm.DB
}

func (r *SpaceRepostirory) Migration() {
	err := r.db.AutoMigrate(&Space{})
	if err != nil {
		fmt.Print(err)
	}
}

func (r *SpaceRepostirory) FindSpace(userid uint) (*Space, error) {
	var space Space
	err := r.db.Where("UserId=?", userid).First(&space).Error
	if err != nil {
		return nil, err
	}
	return &space, nil
}
