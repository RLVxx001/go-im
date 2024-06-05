package userImg

// 用户消息结构体service结构体
type Service struct {
	r Repository
}

// 实例化service
func NewService(r Repository) *Service {
	r.Migration()
	return &Service{
		r: r,
	}
}

// 新增图片
func (s *Service) Create(img *UserImg) error {
	return s.r.Create(img)
}

// 根据用户id查找图片
func (s *Service) GetByUser(userId uint) ([]UserImg, error) {
	imgs, err := s.r.GetByUser(userId)
	if err != nil {
		return nil, ErrFind
	}
	return imgs, nil
}

// 根据图片id查询图片
func (s *Service) GetById(id uint) (*UserImg, error) {
	img, err := s.r.GetById(id)
	if err != nil {
		return nil, ErrFind
	}
	return img, nil
}

// 删除图片
func (s *Service) Delete(id, userId uint) error {
	img, err := s.r.GetById(id)
	if err != nil || img.UserId != userId {
		return ErrDelete
	}
	s.r.Delete(id)
	return nil
}
