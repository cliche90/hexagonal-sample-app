package main

import (
	"hexagonal-sample-app/pkg/core/beer"
	"hexagonal-sample-app/pkg/core/review"
	"hexagonal-sample-app/pkg/http/rest"
	"hexagonal-sample-app/pkg/repository/json"
)

func main() {
	r := &json.Repository{}

	bs := beer.New(r)
	rs := review.New(r)

	s := rest.New(bs, rs)
	s.Start()
	s.WaitStopSignal()
}