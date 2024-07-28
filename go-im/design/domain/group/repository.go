package group

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
	err := r.db.AutoMigrate(&Group{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户-用户对话
func (r *Repository) Create(u *Group) error {
	return r.db.Create(u).Error
}

// 更改操作
func (r *Repository) Update(u *Group) error {
	return r.db.Model(u).Updates(u).Error
}

// 根据id查找
func (r *Repository) GetById(id uint) (*Group, error) {
	var group Group
	err := r.db.Where("ID=?", id).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// 根据公开群号查找
func (r *Repository) GetByGroupId(id string) (*Group, error) {
	var group Group
	err := r.db.Where("GroupId=?", id).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// 查找
func (r *Repository) Fids(u uint) ([]Group, error) {
	var users []Group
	err := r.db.Where("UserOwner=?", u).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 删除群
func (r *Repository) Delete(id uint) error {
	return r.db.Unscoped().Where("ID=?", id).Delete(&Group{}).Error
}

// 更新群头像
func (r *Repository) UpdateImg(img string, id uint) error {
	return r.db.Model(&Group{}).Where("ID=?", id).Update("Img", img).Error
}
