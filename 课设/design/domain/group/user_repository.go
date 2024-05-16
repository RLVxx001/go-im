package group

import (
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

// 实例化
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// 生成表
func (r *UserRepository) Migration() {
	err := r.db.AutoMigrate(&GroupUser{})
	if err != nil {
		log.Print(err)
	}
}

// 新增群用户
func (r *UserRepository) Create(groupUser *GroupUser) error {
	return r.db.Create(groupUser).Error
}

// 更改群用户
func (r *UserRepository) Update(groupUser *GroupUser) error {
	return r.db.Save(groupUser).Error
}

// 查找群用户
func (r *UserRepository) GetGroupUser(id, userid uint) (*GroupUser, error) {
	var groupUser GroupUser
	err := r.db.Where("GroupId=?", id).Where("UserId=?", userid).First(&groupUser).Error
	if err != nil {
		return nil, err
	}
	return &groupUser, nil
}

// 查找群下所有用户
func (r *UserRepository) GetGroupUsers(id uint) ([]GroupUser, error) {
	var groupUser []GroupUser
	err := r.db.Where("GroupId=?", id).Find(&groupUser).Error
	if err != nil {
		return nil, err
	}
	return groupUser, nil
}

// 删除群用户
func (r *UserRepository) Delete(id, userid uint) error {
	return r.db.Where("GroupId=?", id).Where("UserId=?", userid).Delete(&GroupUser{}).Error
}

// 删除群内所有用户
func (r *UserRepository) DeleteAllin(id uint) error {
	return r.db.Where("GroupId=?", id).Delete(&GroupUser{}).Error
}
