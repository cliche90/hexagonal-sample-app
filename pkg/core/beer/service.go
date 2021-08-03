package beer

type Service interface {
	AddBeer(Beer) error
}

type service struct {
	repo   Repository
}

func (s *service) AddBeer(b Beer) error {
	return s.repo.AddBeer(b)
}

// New Dependencies Injection
func New(r Repository) Service {
	return &service{
		repo:   r,
	}
}
