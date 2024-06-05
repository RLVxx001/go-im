package userImg

import "gorm.io/gorm"

type UserImg struct {
	gorm.Model
	Img    string
	UserId uint
}

func NewUserImg(img string, userId uint) *UserImg {
	return &UserImg{
		Img:    img,
		UserId: userId,
	}
}
