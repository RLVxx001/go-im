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
	//校验
	if _, err := c.r.Fid(u.UserOwner, u.UserTarget); err == nil {
		return ErrNotCreate
	}

	if err := c.r.Create(u); err != nil {
		return ErrNotCreate
	}
	return nil
}

// 更改操作
func (c *Service) Update(u *UsertoUser) error {
	if err := c.r.Update(u); err != nil {
		return ErrNotUpdate
	}
	return nil
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
func (c *Service) Revocation(u *UsertoUser, m *UserMessage) error {
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

// 个人删除消息操作（不可逆）
func (c *Service) DeleteMessage(u *UsertoUser, m *UserMessage) error {
	//校验
	if _, err := c.r.Fid(u.UserOwner, u.UserTarget); err != nil {
		return ErrNotDelete
	}

	if u.ID != m.UsertoUserId {
		return ErrNotDelete
	}

	if err := c.messageRepository.Delete(m); err != nil {
		return ErrNotDelete
	}

	return nil
}

// 查找消息
func (c *Service) Fid(u *UsertoUser) (*UsertoUser, error) {
	//校验
	u1, err := c.r.Fid(u.UserOwner, u.UserTarget)
	if err != nil {
		return nil, ErrNotFid
	}
	u1.UserMassages = c.messageRepository.Fid(u1.ID)

	return u1, nil
}
