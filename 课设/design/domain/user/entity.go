package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`
	Account   string `gorm:"type:varchar(30)"`
	Password  string `gorm:"type:varchar(100)"`
	Password2 string `gorm:"-"`
	Salt      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100)"`
	Token     string `gorm:"type:varchar(500)"`
	IsDeleted bool
	IsAdmin   bool
}

func NewUser(username, account, password, password2, email string) *User {
	return &User{
		Username:  username,
		Account:   account,
		Password:  password,
		Password2: password2,
		Email:     email,
		IsAdmin:   false,
		IsDeleted: false,
	}
}
