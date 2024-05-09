package usertoUser

import (
	"log"
)

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
func (c *Service) Create(u *UsertoUser) (*UsertoUser, error) {
	//校验
	us, err := c.r.Fid(u.UserOwner, u.UserTarget)
	if err == nil {
		if us.IsDeleted {
			us.IsDeleted = false
			if err := c.r.Update(us); err != nil {
				return nil, ErrNotCreate
			}
			return us, nil
		}
		return nil, ErrNotCreate
	}

	if err := c.r.Create(u); err != nil {
		return nil, ErrNotCreate
	}
	return u, nil
}

// 更改操作
func (c *Service) Update(u *UsertoUser) error {
	fid, err := c.r.Fid(u.UserOwner, u.UserTarget)
	if err != nil {
		return ErrNotUpdate
	}
	u.ID = fid.ID
	if err := c.r.Update(u); err != nil {
		return ErrNotUpdate
	}
	return nil
}

// 发送消息
func (c *Service) Send(u *UsertoUser, st string) (*UserMessage, *UserMessage, error) {
	u1, err1 := c.r.Fid(u.UserOwner, u.UserTarget)

	if err1 != nil || u1.IsDeleted {
		return nil, nil, ErrNotSend
	}

	u2, err2 := c.r.Fid(u.UserTarget, u.UserOwner)
	if err2 != nil {
		return nil, nil, ErrNotSend
	}

	m := NewUserMessage(u1.ID, st)
	if err := c.messageRepository.Create(m); err != nil { //创建发送者消息
		return nil, nil, ErrNotSend
	}

	m.Key = m.ID
	m.IsRead = true
	if err := c.messageRepository.Update(m); err != nil { //修改发送者消息KEY
		return nil, nil, ErrNotSend
	}
	m1 := NewUserMessage(u2.ID, m.Message)
	m1.Key = m.Key

	if err := c.messageRepository.Create(m1); err != nil { //创建接收者消息
		return nil, nil, ErrNotSend
	}
	return m, m1, nil
}

// 撤回
func (c *Service) Revocation(u *UsertoUser) error {
	//校验
	if _, err := c.r.Fid(u.UserOwner, u.UserTarget); err != nil {
		return ErrNotRevocation
	}

	m := u.UserMassages[0]

	if _, err := c.messageRepository.FidKey(u.ID, m.Key); err != nil {
		return ErrNotRevocation
	}

	if err := c.messageRepository.Deletes(m.Key); err != nil {
		return ErrNotRevocation
	}
	return nil
}

// 个人删除消息操作（不可逆）
func (c *Service) DeleteMessage(u *UsertoUser) error {
	//校验
	if _, err := c.r.Fid(u.UserOwner, u.UserTarget); err != nil {
		return ErrNotDelete
	}

	m := u.UserMassages[0]

	if err := c.messageRepository.Delete(u.ID, m.Key); err != nil {
		return ErrNotDelete
	}

	return nil
}

// 查找消息
func (c *Service) FidMessage(u *UsertoUser) (*UsertoUser, error) {
	//校验
	u1, err := c.r.Fid(u.UserOwner, u.UserTarget)
	if err != nil {
		return nil, ErrNotFid
	}
	u1.UserMassages = c.messageRepository.Fid(u1.ID)

	return u1, nil
}

// 查找用户-用户实体
func (c *Service) Fid(u1, u2 uint) (*UsertoUser, error) {
	us, err := c.r.Fid(u1, u2)
	if err != nil {
		return nil, ErrNotUsers
	}
	return us, nil
}

// 已读消息
func (c *Service) ReadMessage(u *UserMessage) {
	u.IsRead = true
	if err := c.messageRepository.Update(u); err != nil {
		log.Println(err)
	}
}
