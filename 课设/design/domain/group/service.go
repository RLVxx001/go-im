package group

import (
	"log"
	"sort"
)

type Service struct {
	r                 Repository
	messageRepository MessageRepository
	userRepository    UserRepository
}

func NewService(r Repository,
	messageRepository MessageRepository,
	userRepository UserRepository) *Service {

	r.Migration()
	messageRepository.Migration()
	userRepository.Migration()
	return &Service{
		r:                 r,
		messageRepository: messageRepository,
		userRepository:    userRepository,
	}
}

// 创建群
func (s *Service) CreateGroup(group *Group) error {
	//先查询是否存在该群号
	if _, err := s.r.GetByGroupId(group.GroupId); err == nil {
		return ErrGroupId
	}
	group.ID = 0
	if s.r.Create(group) != nil {
		return ErrCreate
	}
	return nil
}

// 通过group的id查询group
func (s *Service) GetById(id uint) (*Group, error) {
	byId, err := s.r.GetById(id)
	if err != nil {
		return nil, ErrNotGroupId
	}
	return byId, nil
}

// 查询群下用户的实体
func (s *Service) GetGroupUser(groupId, userId uint) (*GroupUser, error) {
	groupUser, err := s.userRepository.GetGroupUser(groupId, userId)
	if err != nil {
		return nil, ErrNotGroupUser
	}
	return groupUser, nil
}

// 查找用户账号下所有管理的群聊
func (s *Service) FidMyManage(userId uint) ([]Group, error) {
	users, err := s.userRepository.GetByUserIdGroupUsers(userId)
	if err != nil {
		return nil, ErrFid
	}
	var groups []Group
	for _, i := range users {
		if i.IsAdmin == 0 {
			continue
		}
		group, err := s.r.GetById(i.GroupId)
		if err == nil {
			groupUsers, err := s.userRepository.GetGroupUsers(group.ID)
			if err == nil {
				group.GroupUsers = groupUsers
			}
			group.GroupMessages = s.messageRepository.Fid(group.ID, userId)

			if len(group.GroupMessages) != 0 {
				if group.UpdatedAt.Before(group.GroupMessages[len(group.GroupMessages)-1].UpdatedAt) {
					group.UpdatedAt = group.GroupMessages[len(group.GroupMessages)-1].UpdatedAt
				}
			}
		}
		groups = append(groups, *group)
	}
	sort.Sort(gr(groups))
	return groups, nil
}

// 通过id查找自己的所有群
func (s *Service) FidGroups(userid uint) ([]Group, error) {
	users, err := s.userRepository.GetByUserIdGroupUsers(userid)
	if err != nil {
		return nil, ErrFid
	}
	var groups []Group
	for _, i := range users {
		group, err := s.r.GetById(i.GroupId)
		if err == nil {
			group.GroupMessages = s.messageRepository.Fid(group.ID, userid)
			groupUsers, err := s.userRepository.GetGroupUsers(group.ID)
			if err == nil {
				group.GroupUsers = groupUsers
				mp := make(map[uint]GroupUser)
				for _, o := range groupUsers {
					mp[o.UserId] = o
				}

				for o, k := range group.GroupMessages {
					group.GroupMessages[o].SenderUser = mp[k.MessageSender]
				}
			}
			if len(group.GroupMessages) != 0 {
				if group.UpdatedAt.Before(group.GroupMessages[len(group.GroupMessages)-1].UpdatedAt) {
					group.UpdatedAt = group.GroupMessages[len(group.GroupMessages)-1].UpdatedAt
				}
			}
		}
		groups = append(groups, *group)
	}
	sort.Sort(gr(groups))
	return groups, nil
}

type gr []Group

func (a gr) Len() int           { return len(a) }
func (a gr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a gr) Less(i, j int) bool { return a[i].UpdatedAt.After(a[j].UpdatedAt) }

// 通过id和群号查找群
func (s *Service) FidGroup(id, userid uint) (*Group, error) {
	if _, err := s.userRepository.GetGroupUser(id, userid); err != nil {
		return nil, ErrNotGroupUser
	}
	group, err := s.r.GetById(id)
	if err != nil {
		return nil, ErrNotGroupId
	}
	group.GroupMessages = s.messageRepository.Fid(id, userid)
	group.GroupUsers, err = s.userRepository.GetGroupUsers(id)
	if err != nil {
		return nil, ErrFid
	}
	mp := make(map[uint]GroupUser)
	for _, o := range group.GroupUsers {
		mp[o.UserId] = o
	}
	for o, k := range group.GroupMessages {
		group.GroupMessages[o].SenderUser = mp[k.MessageSender]
	}
	return group, nil
}

// 更改群信息
func (s *Service) UpdateGroup(group *Group, userid uint) error {
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
func (s *Service) DeleteGroup(id, userid uint) error {
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
func (s *Service) UpdateGroupUser(groupUser *GroupUser, userid uint) error {
	//先查群目标用户在这个群的权限
	groupUser1, err1 := s.userRepository.GetGroupUser(groupUser.GroupId, groupUser.UserId)
	if err1 != nil {
		return ErrNotGroupUser
	}

	//如果是修改本人的权限
	if groupUser.UserId == userid {
		groupUser.IsAdmin = groupUser1.IsAdmin //除了权限和禁言不能改以外
		groupUser.IsGag = groupUser1.IsGag
		if s.userRepository.Update(groupUser) != nil {
			return ErrUpdate
		}
		return nil
	}
	//再查该群用户在这个群的权限
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
func (s *Service) DeleteGroupUser(groupUser *GroupUser, userid uint) error {
	//先查群目标用户在这个群的权限
	groupUser1, err1 := s.userRepository.GetGroupUser(groupUser.GroupId, groupUser.UserId)
	if err1 != nil {
		return ErrNotGroupUser
	}
	//如果是自己想要退群
	if groupUser.UserId == userid {
		if s.userRepository.Delete(groupUser.GroupId, groupUser.UserId) != nil {
			return ErrNotDelete
		}
		return nil
	}
	//再查群该用户在这个群的权限
	groupUser2, err2 := s.userRepository.GetGroupUser(groupUser.GroupId, userid)
	if err2 != nil || groupUser1.IsAdmin >= groupUser2.IsAdmin { //如果用户不是群中用户或者权限比目标用户低
		return ErrNotUpdate
	}
	if s.userRepository.Delete(groupUser.GroupId, groupUser.UserId) != nil {
		return ErrNotDelete
	}
	return nil
}

// 新增群用户
func (s *Service) CreateGroupUser(groupId, id, userid uint) error {
	//查询该用户是否已经存在群中
	if _, err := s.userRepository.GetGroupUser(groupId, id); err == nil {
		return ErrNotCreateUser
	}
	user, err := s.userRepository.GetGroupUser(groupId, userid)
	if err != nil || user.IsAdmin == 0 {
		return ErrNotUpdate
	}
	if s.userRepository.Create(NewGroupUser(groupId, id)) != nil {
		return ErrCreate
	}
	return nil
}

// 发送信息
func (s *Service) SendMessage(id, userid uint, message, img string) ([]GroupMessage, error) {
	//查询该用户是否已经存在群中
	groupUser, err := s.userRepository.GetGroupUser(id, userid)
	if err != nil {
		return nil, ErrNotGroupUser
	}
	if groupUser.IsGag {
		return nil, ErrGag
	}
	users, err := s.userRepository.GetGroupUsers(id)
	if err != nil {
		return nil, ErrNotSend
	}
	var messages []GroupMessage
	groupMessage := NewGroupMessage(userid, userid, id, message)
	groupMessage.Img = img
	if s.messageRepository.Create(groupMessage) != nil {
		return nil, ErrNotSend
	}
	groupMessage.MessageKey = groupMessage.ID
	if s.messageRepository.Update(groupMessage) != nil {
		return nil, ErrNotSend
	}
	messages = append(messages, *groupMessage)
	for _, j := range users {
		if j.UserId != userid {
			k := NewGroupMessage(j.UserId, userid, id, message)
			k.Img = img
			k.MessageKey = groupMessage.MessageKey
			if s.messageRepository.Create(k) == nil {
				messages = append(messages, *k)
			}
		}
	}
	return messages, nil
}

// 撤回群内消息
func (s *Service) RevocationMessage(messageId, userid uint) ([]GroupMessage, error) {
	//先根据给的id查询
	groupMessage, err := s.messageRepository.FidId(messageId)
	if err != nil || groupMessage.MessageOwner != userid {
		return nil, ErrRevocation
	}
	//在查询撤回人的身份
	user1, err := s.userRepository.GetGroupUser(groupMessage.GroupId, userid)
	if err != nil {
		return nil, ErrRevocation
	}
	//查询发送人身份
	user2, err := s.userRepository.GetGroupUser(groupMessage.GroupId, groupMessage.MessageSender)
	if err != nil {
		return nil, ErrRevocation
	}
	//如果消息不是自己发的 并且撤回人的权限并不比发送人权限大
	if groupMessage.MessageSender != userid && user2.IsAdmin >= user1.IsAdmin {
		return nil, ErrNotUpdate
	}
	//收集所有群内用户
	groupMessages, err := s.messageRepository.FidKey(groupMessage.GroupId, groupMessage.MessageKey)
	if err != nil {
		return nil, ErrRevocation
	}
	if s.messageRepository.Revocation(groupMessage.MessageKey) != nil {
		return nil, ErrRevocation
	}
	return groupMessages, nil
}

// 删除个人群信息
func (s *Service) DeleteMessage(messageId, userid uint) error {
	//先根据给的id查询
	groupMessage, err := s.messageRepository.FidId(messageId)
	if err != nil || groupMessage.MessageOwner != userid {
		return ErrNotDelete
	}
	if s.messageRepository.Delete(groupMessage.GroupId, groupMessage.MessageKey, userid) != nil {
		return ErrNotDelete
	}
	return nil
}

// 删除个人全部信息
func (s *Service) DeletesMessage(groupId, userid uint) error {
	//先查找该用户是否在群中
	_, err := s.userRepository.GetGroupUser(groupId, userid)
	if err != nil {
		return ErrNotGroupUser
	}
	if s.messageRepository.Deletes(groupId, userid) != nil {
		return ErrNotDelete
	}
	return nil
}

// 已读群消息
func (s *Service) ReadMessage(id, userId uint) {
	err := s.messageRepository.ReadMessage(id, userId)
	if err != nil {
		log.Print(err)
	}
}
