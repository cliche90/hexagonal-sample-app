package review

type Service interface {
	AddReview(Review) error
	GetReviews() ([]Review, error)
}

type service struct {
	repo Repository
}

func (s *service) AddReview(r Review) error {
	return s.repo.AddReview(r)
}

func (s *service) GetReviews() ([]Review, error) {
	return s.repo.GetReviews()
}


func New(r Repository) Service {
	return &service{
		repo: r,
	}
}
