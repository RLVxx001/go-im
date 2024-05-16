package group

import (
	"gorm.io/gorm"
	"log"
)

type MessageRepository struct {
	db *gorm.DB
}

// 实例化
func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

// 生成表
func (r *MessageRepository) Migration() {
	err := r.db.AutoMigrate(&GroupMessage{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户消息
func (r *MessageRepository) Create(u *GroupMessage) error {
	result := r.db.Create(u)
	return result.Error
}

// 删除消息
func (r *MessageRepository) Delete(id, key, userid uint) error {
	tx := r.db.Unscoped().Where("GroupId=?", id).Where("MessageKey=?", key).Where("MessageOwner=?", userid).Delete(&GroupMessage{})
	return tx.Error
}

// 批量删除消息
func (r *MessageRepository) Deletes(id, userid uint) error {
	tx := r.db.Unscoped().Where("GroupId=?", id).Where("MessageOwner=?", userid).Delete(&GroupMessage{})
	return tx.Error
}

// 删除群内所有消息
func (r *MessageRepository) DeleteAllin(id uint) error {
	tx := r.db.Unscoped().Where("GroupId=?", id).Delete(&GroupMessage{})
	return tx.Error
}

// 撤回消息
func (r *MessageRepository) Revocation(key uint) error {
	tx := r.db.Unscoped().Where("MessageKey=?", key).Delete(&GroupMessage{})
	return tx.Error
}

// 查询消息
func (r *MessageRepository) Fid(id, userid uint) []GroupMessage {
	var us []GroupMessage
	r.db.Where("GroupId=?", id).Where("MessageOwner=?", userid).Find(&us)
	return us
}

// 查询消息
func (r *MessageRepository) FidKey(id, key uint) ([]GroupMessage, error) {
	var us []GroupMessage
	err := r.db.Where("GroupId=?", id).Where("MessageKey=?", key).Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

// 根据id查询消息
func (r *MessageRepository) FidId(messageId uint) (GroupMessage, error) {
	var us GroupMessage
	err := r.db.Where("ID=?", messageId).First(&us).Error
	if err != nil {
		return GroupMessage{}, err
	}
	return us, nil
}

// 更改消息
func (r *MessageRepository) Update(u *GroupMessage) error {
	return r.db.Save(u).Error
}

// 已读消息
func (r *MessageRepository) ReadMessage(id, userid uint) error {
	return r.db.Model(&GroupMessage{}).
		Where("IsRead=?", false).
		Where("GroupId=?", id).
		Where("MessageOwner=?", userid).
		Update("IsRead", true).
		Error
}
