package space

import (
	"fmt"
	"gorm.io/gorm"
)

type SpaceRepository struct {
	db *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) *SpaceRepository {
	return &SpaceRepository{
		db: db,
	}
}

func (r *SpaceRepository) Migration() {
	err := r.db.AutoMigrate(&Space{})
	if err != nil {
		fmt.Print(err)
	}
}

func (r *SpaceRepository) Create(userId uint) error {
	var space Space
	space.UserId = userId
	space.EnterPermit = false
	var trends []SpaceTrends
	space.SpaceTrends = trends
	return r.db.Create(&space).Error
}

func (r *SpaceRepository) Find(userid uint) (Space, error) {
	var space Space
	err := r.db.Where("UserId=?", userid).First(&space).Error
	if err != nil {
		print(err)
	}
	return space, err
}
