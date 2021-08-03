package beer

type Service interface {
	AddBeer(Beer) error
	GetBeers() ([]Beer, error)
}

type service struct {
	repo Repository
}

func (s *service) AddBeer(b Beer) error {
	return s.repo.AddBeer(b)
}

func (s *service) GetBeers() ([]Beer, error) {
	return s.repo.GetBeers()
}

func New(r Repository) Service {
	return &service{
		repo: r,
	}
}
