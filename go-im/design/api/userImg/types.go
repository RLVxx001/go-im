package userImg

import "design/domain/userImg"

type Img struct {
	Id     uint   `json:"id"`
	Img    string `json:"url"`
	UserId uint   `json:"userId"`
}

func ToImg(img userImg.UserImg) Img {
	return Img{
		Id:     img.ID,
		Img:    img.Img,
		UserId: img.UserId,
	}
}
