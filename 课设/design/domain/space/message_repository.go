package space

import (
	"design/domain/user"
	"fmt"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) Migration() {
	err := r.db.AutoMigrate(&Message{})
	if err != nil {
		fmt.Print(err)
	}
}

func (r *MessageRepository) Create(message Message) error {
	return r.db.Create(&message).Error
}

func (r *MessageRepository) Find(messageId uint) Message {
	var message Message
	r.db.Where("ID=?", messageId).First(&message)
	return message
}

func (r *MessageRepository) Finds(spaceId uint) []Message {
	var message []Message
	r.db.Where("SpaceId=?", spaceId).Find(&message)
	return message
}

func (r *MessageRepository) FindUser(userId uint) user.User {
	var user user.User
	r.db.Where("ID=?", userId).First(&user)
	return user
}

func (r *MessageRepository) Delete(messageId uint) error {
	return r.db.Where("ID=?", messageId).Delete(&Message{}).Error
}
