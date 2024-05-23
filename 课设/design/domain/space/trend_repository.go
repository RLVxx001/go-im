package space

import (
	"gorm.io/gorm"
	"log"
)

type TrendsRepository struct {
	db *gorm.DB
}

func NewTrendsRepository(db *gorm.DB) *TrendsRepository {
	return &TrendsRepository{
		db: db,
	}
}

func (r *TrendsRepository) Migration() {
	err := r.db.AutoMigrate(&SpaceTrends{})
	if err != nil {
		log.Print(err)
	}
}

// create
func (r *TrendsRepository) Create(trends *SpaceTrends) error {
	return r.db.Create(&trends).Error
}

func (r *TrendsRepository) Delete(trendid uint) error {
	return r.db.Where("ID=?", trendid).Delete(SpaceTrends{}).Error
}

func (r *TrendsRepository) Find(trendId uint) (SpaceTrends, error) {
	var trend SpaceTrends
	err := r.db.Where("ID=?", trendId).First(&trend).Error
	return trend, err
}

func (r *TrendsRepository) Update(trend SpaceTrends) error {
	return r.db.Where("ID=?", trend.ID).Update("Comment", trend).Error
}
