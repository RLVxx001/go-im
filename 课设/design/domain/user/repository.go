package user

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// 实例化
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户
func (r *Repository) Create(u *User) error {
	result := r.db.Create(u)
	return result.Error
}

// 根据用户名查询用户
func (r *Repository) GetByName(name string) (User, error) {
	var user User
	err := r.db.Where("UserName=?", name).Where("IsDeleted=?", 0).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// 根据邮件查询用户
func (r *Repository) GetByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("Email=?", email).Where("IsDeleted=?", 0).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// 根据昵称查询用户
func (r *Repository) GetByAccount(account string) ([]User, error) {
	var users []User
	err := r.db.Where("Account=?", account).Where("IsDeleted=?", 0).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 添加测试数据
func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "admin", "admin", "admin", "admin@admin")

	user.IsAdmin = true
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username,
			Password: user.Password,
		}).FirstOrCreate(&user)

	user = NewUser("user", "user", "user", "user", "user")
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username,
			Password: user.Password,
		}).FirstOrCreate(&user)
}

// 更新用户
func (r *Repository) Update(u *User) error {
	return r.db.Save(&u).Error
}
