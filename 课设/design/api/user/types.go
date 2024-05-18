package user

import "design/domain/user"

// 创建用户请求结构体
type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Email     string `json:"email"`
}

// 创建用户响应
type CreateUserResponse struct {
	Username string `json:"username"`
}

// 登录请求
type LoginRequest struct {
	Text     string `json:"username"` //可是邮箱也可也是账号
	Password string `json:"password"`
}

// 登录响应
type LoginResponse struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	UserId   uint   `json:"userId"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Img      string `json:"img"`
	Signed   string `json:"signed"`   //个性签名
	Birthday string `json:"birthday"` //出生日期
}

func ToLoginResponse(currentUser user.User) LoginResponse {
	return LoginResponse{
		Username: currentUser.Username,
		UserId:   currentUser.ID,
		Token:    currentUser.Token,
		Account:  currentUser.Account,
		Email:    currentUser.Email,
		Img:      currentUser.Img,
		Signed:   currentUser.Signed,   //个性签名
		Birthday: currentUser.Birthday, //出生日期
	}
}
