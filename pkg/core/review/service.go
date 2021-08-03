package review

type Service interface {
	AddReview(Review) error
}

type service struct {
	repo Repository
}

func (s *service) AddReview(r Review) error {
	return nil
}

func New(r Repository) Service {
	return &service{
		repo: r,
	}
}
