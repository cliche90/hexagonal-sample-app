package json

import (
	"hexagonal-sample-app/pkg/core/beer"
	"hexagonal-sample-app/pkg/core/review"
)

type Repository struct {
}

func (r *Repository) AddBeer(b beer.Beer) error {
	return nil
}

func (r *Repository) AddReview(review review.Review) error {
	return nil
}

func NewRepository() *Repository {
	return &Repository{

	}
}