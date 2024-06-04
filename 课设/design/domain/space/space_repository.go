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
		return Space{}, err
	}
	r.db.Where("UserId=?", userid).Find(&space.SpaceTrends)
	for i := 0; i < len(space.SpaceTrends); i++ {
		r.db.Where("ID=?", userid).First(&space.SpaceTrends[i].User)
	}
	if err != nil {
		print(err)
	}
	return space, err
}

func (r *SpaceRepository) Update(space Space) error {
	return r.db.Where("ID=?", space.ID).Update("SpaceTrends", space.SpaceTrends).Error
}
