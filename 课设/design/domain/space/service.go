package space

import (
	"fmt"
	"gorm.io/gorm"
)

type Service struct {
	space   SpaceRepository
	trend   TrendsRepository
	comment CommentRepository
}

func NewService(r SpaceRepository, t TrendsRepository, c CommentRepository) *Service {
	r.Migration()
	t.Migration()
	c.Migration()
	//r.InsertSampleData()
	return &Service{
		space:   r,
		trend:   t,
		comment: c,
	}
}

func (s *Service) CreateSpace(userId uint) error {
	//var user uint
	_, err := s.space.Find(userId)
	if err == nil {
		return nil
	}
	err = s.space.Create(userId)
	if err != nil {
		print(err)
	}
	return err
}

func (r *Service) FindSpace(userid uint) (Space, error) {
	var space Space
	space, err := r.space.Find(userid)
	if err != nil {
		return space, err
	}
	return space, nil
}

func (r *Service) FindTrend(trendId uint) (SpaceTrends, error) {
	trend, err := r.trend.Find(trendId)
	comments, err := r.comment.Find(trendId)
	if err != nil {
		print(err)
	}
	trend.Comments = comments
	return trend, err
}

func (r *Service) FindTrends(userid uint) ([]SpaceTrends, error) {
	space, err := r.space.Find(userid)
	if err != nil {
		print(err)
	}
	fmt.Printf("%v", space.SpaceTrends)
	return space.SpaceTrends, err
}

func (r *Service) DeleteTrends(trend SpaceTrends) error {
	err := r.trend.Delete(trend.ID)
	if err != nil {
		print(err)
	}
	return err
}

func (r *Service) CreateTrends(trend SpaceTrends) error {
	err := r.trend.Create(trend)
	space, err := r.space.Find(trend.UserId)
	space.SpaceTrends = append(space.SpaceTrends, trend)
	fmt.Printf("%v", space)
	r.space.Update(space)
	if err != nil {
		print(err)
	}
	return err
}

func (r *Service) FindComments(trendId uint) ([]Comment, error) {
	comments, err := r.comment.Find(trendId)
	return comments, err
}

func (r *Service) CreateComment(userId uint, detail string, trendId uint) error {
	comment := r.comment.Create(Comment{
		Model:         gorm.Model{},
		UserId:        userId,
		Praise:        1,
		Content:       detail,
		TrendsId:      trendId,
		ToUserId:      userId,
		SpaceTrendsId: trendId,
	})
	fmt.Printf("%v\n", comment)
	trend, err := r.trend.Find(comment.TrendsId)
	trend.Comments = append(trend.Comments, comment)
	err = r.trend.Update(trend)
	if err != nil {
		print(err)
	}
	return err
}

func (r *Service) AddComment(comment Comment) error {
	var trend SpaceTrends
	var err error
	trend, err = r.trend.Find(comment.TrendsId)
	if err != nil {
		print(err)
	}
	var tmp []Comment
	tmp = trend.Comments
	trend.Comments = tmp
	err = r.trend.Update(trend)
	if err != nil {
		print(err)
	}
	return err
}

func (r *Service) DeleteComment(comment Comment) error {
	err := r.comment.Delete(comment.ID)
	if err != nil {
		print(err)
	}
	return err
}
