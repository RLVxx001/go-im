package userApplication

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
	//判断对方是否已经发送申请请求
	if userTO, err := c.r.Fid(u.UserTarget, u.UserOwner); err == nil {
		userTO.IsAccept = true
		if err := c.r.Update(userTO); err != nil {
			return nil, ErrApplication
		}

		u.IsAccept = true
		return u, nil
	}

	//校验
	us, err := c.r.Fid(u.UserOwner, u.UserTarget)
	if err == nil { //已经发送过则覆盖申请
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
func (c *Service) Update(u1, u2 uint) error {
	fid, err := c.r.Fid(u1, u2)
	if err == nil { //如果之前申请过便删除
		if err := c.r.Delete(fid.ID); err != nil {
			return ErrNotUpdate
		}
		return nil
	}

	fid, err = c.r.Fid(u2, u1)
	if err != nil {
		return ErrNotUpdate
	}
	fid.IsDown = true
	if err := c.r.Update(fid); err != nil {
		return ErrNotUpdate
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
