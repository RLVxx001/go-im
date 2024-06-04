package space

import "gorm.io/gorm"

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(comment Comment) error {
	return r.db.Create(&comment).Error
}

func (r *CommentRepository) Migration() {
	err := r.db.AutoMigrate(&Comment{})
	if err != nil {
		print(err)
	}
}

func (r *CommentRepository) Delete(commentId uint) error {
	return r.db.Unscoped().Where("ID=?", commentId).Delete(&Comment{}).Error
}

func (r *CommentRepository) Find(trendId uint) ([]Comment, error) {
	var comments []Comment
	err := r.db.Where("TrendsId=?", trendId).Find(&comments).Error
	for i := 0; i < len(comments); i++ {
		r.db.Where("ID=?", comments[i].UserId).First(&comments[i].User)
	}
	return comments, err
}
