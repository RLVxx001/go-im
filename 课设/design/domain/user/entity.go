package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`  //用户名
	Account   string `gorm:"type:varchar(30)"`  //账号名
	Password  string `gorm:"type:varchar(100)"` //密码
	Password2 string `gorm:"-"`                 //验证密码（不计入表）
	Salt      string `gorm:"type:varchar(100)"` //token密钥
	Email     string `gorm:"type:varchar(100)"` //邮件
	Token     string `gorm:"type:varchar(500)"` //token
	Img       string `gorm:"type:varchar(500)"` //头像地址
	Signed    string //个性签名
	Birthday  string //出生日期
	IsDeleted bool   //是否被删除
	IsAdmin   bool   //是否是管理
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
