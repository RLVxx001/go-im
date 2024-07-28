package userImg

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// 实例化
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&UserImg{})
	if err != nil {
		log.Print(err)
	}
}

// 创建图片
func (r *Repository) Create(img *UserImg) error {
	return r.db.Create(img).Error
}

// 根据用户查找图片墙
func (r *Repository) GetByUser(userId uint) ([]UserImg, error) {
	var imgs []UserImg
	err := r.db.Where("UserId=?", userId).Find(&imgs).Error
	if err != nil {
		return nil, err
	}
	return imgs, nil
}

// 根据id查找图片
func (r *Repository) GetById(id uint) (*UserImg, error) {
	var img UserImg
	err := r.db.Where("ID=?", id).First(&img).Error
	if err != nil {
		return nil, err
	}
	return &img, err
}

// 删除图片
func (r *Repository) Delete(id uint) {
	r.db.Unscoped().Where("ID=?", id).Delete(&UserImg{})
}
