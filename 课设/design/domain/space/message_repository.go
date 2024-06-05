package space

import (
	"fmt"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func (r *MessageRepository) Migration() {
	err := r.db.AutoMigrate(&Space{})
	if err != nil {
		fmt.Print(err)
	}
}
