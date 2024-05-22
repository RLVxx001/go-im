package space

type Service struct {
	space SpaceRepostirory
	trend TrendsRepository
}

func (r *Service) FindTrends(userid uint) []SpaceTrends {
	space, err := r.space.FindSpace(userid)
	if err != nil {
		print(err)
	}
	return space.SpaceTrends
}

func (r *Service) DeleteTrends(userid uint) {

}
