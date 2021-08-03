package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"hexagonal-sample-app/pkg/core/beer"
	"hexagonal-sample-app/pkg/core/review"
)

type Repository struct {
	beersFilePath   string
	reviewsFilePath string
}

func (r *Repository) AddBeer(b beer.Beer) error {
	bytes, err := ioutil.ReadFile(r.beersFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	beers := beer.Beers{
		Beers: make([]beer.Beer, 0),
	}
	if err := json.Unmarshal(bytes, &beers); err != nil {
		fmt.Println(err.Error())
		return err
	}

	beers.Beers = append(beers.Beers, b)
	beersBytes, _ := json.MarshalIndent(beers, "", "  ")
	_ = ioutil.WriteFile(r.beersFilePath, beersBytes, 0644)
	return nil
}

func (r *Repository) GetBeers() ([]beer.Beer, error) {
	beersBytes, err := ioutil.ReadFile(r.beersFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	beers := beer.Beers{
		Beers: make([]beer.Beer, 0),
	}
	if err := json.Unmarshal(beersBytes, &beers); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return beers.Beers, nil
}

func (r *Repository) AddReview(rv review.Review) error {
	bytes, err := ioutil.ReadFile(r.reviewsFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	reviews := review.Reviews{
		Reviews: make([]review.Review, 0),
	}
	if err := json.Unmarshal(bytes, &reviews); err != nil {
		fmt.Println(err.Error())
		return err
	}

	reviews.Reviews = append(reviews.Reviews, rv)
	reviewsBytes, _ := json.MarshalIndent(reviews, "", "  ")
	_ = ioutil.WriteFile(r.reviewsFilePath, reviewsBytes, 0644)
	return nil
}

func (r *Repository) GetReviews() ([]review.Review, error) {
	reviewsBytes, err := ioutil.ReadFile(r.reviewsFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	reviews := review.Reviews{
		Reviews: make([]review.Review, 0),
	}
	if err := json.Unmarshal(reviewsBytes, &reviews); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return reviews.Reviews, nil
}

func NewRepository(beersFilePath, reviewsFilePath string) *Repository {
	return &Repository{
		beersFilePath: beersFilePath,
		reviewsFilePath: reviewsFilePath,
	}
}
