package beer

type Repository interface {
	AddBeer(Beer) error
	GetBeers() ([]Beer, error)
}
