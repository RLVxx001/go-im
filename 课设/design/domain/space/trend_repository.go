package space

import (
	"gorm.io/gorm"
	"log"
)

type TrendsRepository struct {
	db *gorm.DB
}

func (r *TrendsRepository) Migration() {
	err := r.db.AutoMigrate(&SpaceTrends{})
	if err != nil {
		log.Print(err)
	}
}

// create
func (r *TrendsRepository) CreateTrends(trends SpaceTrends) error {
	return r.db.Create(&trends).Error
}

func (r *TrendsRepository) DeleteTrends(trendid uint) error {
	return r.db.Where("ID=?", trendid).Delete(SpaceTrends{}).Error
}
