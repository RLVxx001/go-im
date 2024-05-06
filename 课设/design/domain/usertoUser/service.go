package usertoUser

// 用户消息结构体service结构体
type Service struct {
	r                 Repository
	messageRepository MessageRepository
}

// 实例化service
func NewUserService(r Repository, messageRepository MessageRepository) *Service {
	r.Migration()
	messageRepository.Migration()
	return &Service{
		r:                 r,
		messageRepository: messageRepository,
	}
}

// 创建用户-用户链接
func (c *Service) Create(u *UsertoUser) error {
	return c.r.Create(u)
}

// 更改操作
func (c *Service) Update(u *UsertoUser) error {
	return c.r.Update(u)
}

// 发送消息
func (c *Service) Send(u *UsertoUser, m *UserMessage) error {
	if u.IsDeleted {
		return ErrNotUser
	}
	u1, err1 := c.r.Fid(u.UserOwner, u.UserTarget)
	if err1 != nil {
		return ErrNotSend
	}
	u2, err2 := c.r.Fid(u.UserTarget, u.UserOwner)
	if err2 != nil {
		return ErrNotSend
	}
	m.UsertoUserId = u1.ID
	if err := c.messageRepository.Create(m); err != nil { //创建发送者消息
		return ErrNotSend
	}
	m.Key = m.ID
	if err := c.messageRepository.Update(m); err != nil { //修改发送者消息KEY
		return ErrNotSend
	}
	m.UsertoUserId = u2.ID
	if err := c.messageRepository.Create(m); err != nil { //创建接收者消息
		return ErrNotSend
	}
	return nil
}

// 撤回
func (c *Service) revocation(u *UsertoUser, m *UserMessage) error {
	//校验
	if _, err := c.r.Fid(u.UserOwner, u.UserTarget); err != nil {
		return ErrNotRevocation
	}

	if u.ID != m.UsertoUserId {
		return ErrNotRevocation
	}

	if err := c.messageRepository.Deletes(m.Key); err != nil {
		return ErrNotRevocation
	}
	return nil
}
