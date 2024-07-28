package user

import (
	"design/utils/hash"
)

// 用户结构体service结构体
type Service struct {
	r Repository
}

// 实例化service
func NewService(r Repository) *Service {
	r.Migration()
	//r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 创建用户
func (c *Service) Create(user *User) error {
	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}
	//用户名存在
	_, err := c.r.GetByName(user.Username)
	if err == nil {
		return ErrUserExistWithName
	}

	//Email存在
	_, err = c.r.GetByEmail(user.Email)
	if err == nil {
		return ErrUserExistWithEmail
	}

	//无效用户名
	if ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}

	//无效密码
	if ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}

	//创建用户
	err = c.r.Create(user)
	return err
}

// email验证用户
func (c *Service) CheckEmailUser(email, password string) (User, error) {
	user, err := c.r.GetByEmail(email)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// id查找用户
func (c *Service) GetById(id uint) (User, error) {
	user, err := c.r.GetById(id)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// email查找用户
func (c *Service) GetEmailUser(email string) (User, error) {
	user, err := c.r.GetByEmail(email)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// 验证用户
func (c *Service) CheckUser(text, password string) (User, error) {
	user1, err := c.CheckUserName(text, password)
	if err == nil {
		return user1, nil
	}
	user2, err := c.CheckUserEmail(text, password)
	if err == nil {
		return user2, nil
	}
	return User{}, ErrUserNotFound
}

// 用户名验证用户
func (c *Service) CheckUserName(name, password string) (User, error) {
	user, err := c.r.GetByName(name)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// Email验证用户
func (c *Service) CheckUserEmail(name, password string) (User, error) {
	user, err := c.r.GetByEmail(name)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// 用户名查找用户
func (c *Service) GetUser(name string) (User, error) {
	user, err := c.r.GetByName(name)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// 昵称查找用户
func (c *Service) GetUserList(account string) (User, error) {
	user, err := c.r.GetByName(account)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// 更新用户
func (c *Service) UpdateUser(user *User) error {
	return c.r.Update(user)
}

// 上传头像
func (c *Service) UpdateImg(img string, id uint) error {
	return c.r.UpdateImg(img, id)
}
