package user

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
	Text     string `json:"username"`
	Password string `json:"password"`
}

// 登录响应
type LoginResponse struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	UserId   uint   `json:"userId"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
