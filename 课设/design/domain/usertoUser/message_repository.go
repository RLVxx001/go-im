package usertoUser

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
	err := r.db.AutoMigrate(&UserMessage{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户消息
func (r *MessageRepository) Create(u *UserMessage) error {
	result := r.db.Create(u)
	return result.Error
}

// 删除消息
func (r *MessageRepository) Delete(utouid, key uint) error {
	tx := r.db.Where("UsertoUserId=?", utouid).Where("`Key`=?", key).Delete(&UserMessage{})
	return tx.Error
}

// 撤回消息
func (r *MessageRepository) Deletes(key uint) error {
	tx := r.db.Where("`Key`=?", key).Delete(&UserMessage{})
	return tx.Error
}

// 查询消息
func (r *MessageRepository) Fid(utouid uint) []UserMessage {
	var us []UserMessage
	r.db.Where("UsertoUserId=?", utouid).Find(&us)
	return us
}

// 查询消息
func (r *MessageRepository) FidKey(utouid, key uint) (UserMessage, error) {
	var us UserMessage
	err := r.db.Where("UsertoUserId=?", utouid).Where("`Key`=?", key).First(&us).Error
	if err != nil {
		return UserMessage{}, err
	}
	return us, nil
}

// 更改消息
func (r *MessageRepository) Update(u *UserMessage) error {
	return r.db.Save(u).Error
}
