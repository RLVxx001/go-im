package userApplication

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
	err := r.db.AutoMigrate(&UserApplication{})
	if err != nil {
		log.Print(err)
	}
}

// 创建申请表
func (r *Repository) Create(u *UserApplication) error {
	return r.db.Create(u).Error
}

// 更改操作
func (r *Repository) Update(u *UserApplication) error {
	return r.db.Omit("CreatedAt").Save(u).Error
}

// 查找
func (r *Repository) Fid(u, tou uint) (*UserApplication, error) {
	var userApplication UserApplication
	err := r.db.Where("UserOwner=?", u).Where("UserTarget=?", tou).Where("IsDown", false).Where("IsAccept", false).First(&userApplication).Error
	if err != nil {
		return nil, err
	}
	return &userApplication, nil
}

// 查找
func (r *Repository) Fids(u uint) ([]UserApplication, error) {
	var users []UserApplication
	err := r.db.Where("UserOwner=?", u).Or("UserTarget = ?", u).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 删除申请操作
func (r *Repository) Delete(id uint) error {
	return r.db.Unscoped().Where("ID=?", id).Delete(&UserApplication{}).Error
}
