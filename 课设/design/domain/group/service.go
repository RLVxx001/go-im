package group

type Server struct {
	r                 *Repository
	messageRepository *MessageRepository
	userRepository    *UserRepository
}

func NewServer(r *Repository,
	messageRepository *MessageRepository,
	userRepository *UserRepository) *Server {

	r.Migration()
	messageRepository.Migration()
	userRepository.Migration()
	return &Server{
		r:                 r,
		messageRepository: messageRepository,
		userRepository:    userRepository,
	}
}

// 创建群
func (s *Server) CreateGroup(group *Group) error {
	//先查询是否存在该群号
	if _, err := s.r.GetByGroupId(group.GroupId); err != nil {
		return ErrGroupId
	}
	group.ID = 0
	if s.r.Create(group) != nil {
		return ErrCreate
	}
	return nil
}

// 更改群信息
func (s *Server) UpdateGroup(group *Group, userid uint) error {
	//先查群该用户在这个群的权限
	groupUser, err := s.userRepository.GetGroupUser(group.ID, userid)
	if err != nil || groupUser.IsAdmin == 0 {
		return ErrNotUpdate
	}
	if s.r.Update(group) != nil {
		return ErrUpdate
	}
	return nil
}

// 删除群
func (s *Server) DeleteGroup(id, userid uint) error {
	//先查群该用户在这个群的权限
	groupUser, err := s.userRepository.GetGroupUser(id, userid)
	if err != nil || groupUser.IsAdmin != 2 {
		return ErrNotDelete
	}
	//先删除所有群消息
	if err = s.messageRepository.DeleteAllin(id); err != nil {
		return ErrNotDelete
	}
	//在删除所有绑定群用户
	if err := s.userRepository.DeleteAllin(id); err != nil {
		return ErrNotDelete
	}
	//最后删除群
	if err := s.r.Delete(id); err != nil {
		return ErrNotDelete
	}
	return nil
}

// 更改群用户信息包括设置管理员
func (s *Server) UpdateGroupUser(groupUser *GroupUser, userid uint) error {
	//先查群目标用户在这个群的权限
	groupUser1, err1 := s.userRepository.GetGroupUser(groupUser.GroupId, groupUser.UserId)
	if err1 != nil {
		return ErrNotGroupUser
	}
	//再查群该用户在这个群的权限
	groupUser2, err2 := s.userRepository.GetGroupUser(groupUser.GroupId, userid)
	if err2 != nil || (groupUser1.IsAdmin >= groupUser2.IsAdmin || groupUser.IsAdmin >= groupUser2.IsAdmin) { //如果用户不是群中用户或者权限比目标用户低
		return ErrNotUpdate
	}
	groupUser.ID = groupUser1.ID
	if s.userRepository.Update(groupUser) != nil {
		return ErrUpdate
	}
	return nil
}

// 删除群用户
func (s *Server) DeleteGroupUser(groupUser *GroupUser, userid uint) error {
	//先查群目标用户在这个群的权限
	groupUser1, err1 := s.userRepository.GetGroupUser(groupUser.GroupId, groupUser.UserId)
	if err1 != nil {
		return ErrNotGroupUser
	}
	//再查群该用户在这个群的权限
	groupUser2, err2 := s.userRepository.GetGroupUser(groupUser.GroupId, userid)
	if err2 != nil || groupUser1.IsAdmin >= groupUser2.IsAdmin { //如果用户不是群中用户或者权限比目标用户低
		return ErrNotUpdate
	}
	if s.userRepository.Delete(groupUser.ID, groupUser.UserId) != nil {
		return ErrNotDelete
	}
	return nil
}

// 新增群用户
func (s *Server) CreateGroupUser(id, userid uint) error {
	//查询该用户是否已经存在群中
	if _, err := s.userRepository.GetGroupUser(id, userid); err == nil {
		return ErrNotCreateUser
	}
	if s.userRepository.Create(NewGroupUser(id, userid)) != nil {
		return ErrCreate
	}
	return nil
}

// 发送信息
func (s *Server) SendMessage(id, userid uint, message string) error {
	//查询该用户是否已经存在群中
	groupUser, err := s.userRepository.GetGroupUser(id, userid)
	if err != nil {
		return ErrNotGroupUser
	}
	if groupUser.IsGag {
		return ErrGag
	}
	return nil
}
