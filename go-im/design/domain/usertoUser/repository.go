package usertoUser

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
	err := r.db.AutoMigrate(&UsertoUser{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户-用户对话
func (r *Repository) Create(u *UsertoUser) error {
	return r.db.Create(u).Error
}

// 更改操作
func (r *Repository) Update(u *UsertoUser) error {
	return r.db.Omit("CreatedAt").Save(u).Error
}

// 删除好友操作
func (r *Repository) DeleteUser(u *UsertoUser) error {
	return r.db.Model(&UsertoUser{}).
		Where("(UserOwner = ? AND UserTarget = ?)", u.UserOwner, u.UserTarget).
		Or("(UserOwner = ? AND UserTarget = ?)", u.UserTarget, u.UserOwner).
		Update("IsDeleted", true).Error
}

// 查找
func (r *Repository) Fid(u, tou uint) (*UsertoUser, error) {
	var usertoUser UsertoUser
	err := r.db.Where("UserOwner=?", u).Where("UserTarget=?", tou).First(&usertoUser).Error
	if err != nil {
		return nil, err
	}
	return &usertoUser, nil
}

// 查找
func (r *Repository) Fids(u uint) ([]UsertoUser, error) {
	var users []UsertoUser
	err := r.db.Where("UserOwner=?", u).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
