package userApplication

import (
	"time"
)

// 用户消息结构体service结构体
type Service struct {
	r Repository
}

// 实例化service
func NewService(r Repository) *Service {
	r.Migration()
	return &Service{
		r: r,
	}
}

// 创建申请表
func (c *Service) Create(u *UserApplication) (*UserApplication, error) {
	//校验
	us, err := c.r.Fid(u.UserOwner, u.Class, u.Target)
	u.FailureTime = time.Now()
	u.FailureTime = u.FailureTime.Add(time.Hour * 24 * 3) //更新失效时间
	if err == nil {                                       //已经发送过则覆盖申请
		u.ID = us.ID
		if err := c.r.Update(u); err != nil {
			return nil, ErrApplication
		}
		return u, nil
	}
	if err := c.r.Create(u); err != nil {
		return nil, ErrApplication
	}
	return u, nil
}

// 拒绝操作
func (c *Service) Refuse(userApplication *UserApplication) error {
	fid, err := c.r.Fid(userApplication.UserOwner, userApplication.Class, userApplication.Target)
	if err != nil {
		return ErrNotUpdate
	}
	fid.Stats = 1
	fid.InviteUser = userApplication.InviteUser

	if err := c.r.Update(fid); err != nil {
		return ErrNotUpdate
	}
	return nil
}

// 接受操作
func (c *Service) Accept(userApplication *UserApplication) error {
	fid, err := c.r.Fid(userApplication.UserOwner, userApplication.Class, userApplication.Target) //先查询是否有此请求
	if err != nil {
		return ErrAccept
	}
	fid.Stats = 2
	fid.InviteUser = userApplication.InviteUser
	if err := c.r.Update(fid); err != nil {
		return ErrAccept
	}
	return nil
}

// 查询单用户下的所有信息
func (c *Service) Fids(userid uint) ([]UserApplication, error) {
	users, err := c.r.Fids(userid)
	if err != nil {
		return nil, ErrNotFid
	}
	return users, nil
}
