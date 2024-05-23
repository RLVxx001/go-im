package space

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
	return &Service{}
}

func (r *Service) FindSpace(userid uint) (Space, error) {
	var space Space
	space, err := r.space.Find(userid)
	if err != nil {
		return space, err
	}
	return space, nil
}

func (r *Service) FindTrends(userid uint) ([]SpaceTrends, error) {
	space, err := r.space.Find(userid)
	if err != nil {
		print(err)
	}
	return space.SpaceTrends, err
}

func (r *Service) DeleteTrends(trend SpaceTrends) error {
	err := r.trend.Delete(trend.ID)
	if err != nil {
		print(err)
	}
	return err
}

func (r *Service) CreateTrends(trend *SpaceTrends) error {
	err := r.trend.Create(trend)
	if err != nil {
		print(err)
	}
	return err
}

func (r *Service) CreateComment(comment Comment) error {
	err := r.comment.Create(comment)
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
