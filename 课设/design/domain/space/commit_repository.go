package space

import "gorm.io/gorm"

type CommentRepository struct {
	db *gorm.DB
}

func (r *CommentRepository) CreateComment(comment Comment) error {
	return r.db.Create(&comment).Error
}

func (r *CommentRepository) Migration() {
	err := r.db.AutoMigrate(&Comment{})
	if err != nil {
		print(err)
	}
}

func (r *CommentRepository) DeleteComment(commentid uint) error {
	return r.db.Unscoped().Where("ID=?", commentid).Delete(&Comment{}).Error
}
