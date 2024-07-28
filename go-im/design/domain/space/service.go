package space

import (
	"fmt"
	"gorm.io/gorm"
)

type Service struct {
	space   SpaceRepository
	trend   TrendsRepository
	comment CommentRepository
	message MessageRepository
}

func NewService(r SpaceRepository, t TrendsRepository, c CommentRepository, m MessageRepository) *Service {
	r.Migration()
	t.Migration()
	c.Migration()
	m.Migration()
	//r.InsertSampleData()
	return &Service{
		space:   r,
		trend:   t,
		comment: c,
		message: m,
	}
}

func (s *Service) DeleteMessage(messageId uint) error {
	return s.message.Delete(messageId)
}

func (s *Service) CreateMessage(userId uint, detail string) Message {
	var message Message
	space := s.message.FindSpace(userId)
	message.SpaceId = space.ID
	message.UserId = userId
	message.Detail = detail
	err := s.message.Create(message)
	if err != nil {
		print(err)
	}
	return message
}

func (s *Service) FindMessage(userId uint) []Message {
	var messages []Message
	space := s.message.FindSpace(userId)
	messages = s.message.Finds(space.ID)
	for i := 0; i < len(messages); i++ {
		messages[i].User = s.message.FindUser(messages[i].UserId)
	}
	return messages
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
	r.comment.Create(Comment{
		Model:         gorm.Model{},
		UserId:        userId,
		Praise:        1,
		Content:       detail,
		TrendsId:      trendId,
		ToUserId:      userId,
		SpaceTrendsId: trendId,
	})
	return nil
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
