package main

import (
	"hexagonal-sample-app/pkg/core/beer"
	"hexagonal-sample-app/pkg/core/review"
	"hexagonal-sample-app/pkg/proto/http"
	"hexagonal-sample-app/pkg/repository/json"
)

func main() {
	r := json.NewRepository(
		"pkg/repository/json/beers.json",
		"pkg/repository/json/reviews.json",
	)

	bs := beer.New(r)
	rs := review.New(r)

	s := http.New(bs, rs)

	s.Start()
	s.WaitStopSignal()
}