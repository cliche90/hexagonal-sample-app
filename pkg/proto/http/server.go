package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"hexagonal-sample-app/pkg/core/beer"
	"hexagonal-sample-app/pkg/core/review"
)

type server struct {
	server *http.Server
	done   chan os.Signal
}

func (s *server) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")
}

func (s *server) WaitStopSignal() {
	<-s.done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}

func addBeer(bs beer.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var b beer.Beer
		_ = json.NewDecoder(req.Body).Decode(&b)
		_ = bs.AddBeer(b)
		_, _ = resp.Write([]byte("beers are added."))
	}
}

func getBeers(bs beer.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		beers, _ := bs.GetBeers()
		beersBytes, _ := json.MarshalIndent(beers, "", "  ")
		_, _ = resp.Write(beersBytes)
	}
}

func addReview(rs review.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var r review.Review
		_ = json.NewDecoder(req.Body).Decode(&r)
		_ = rs.AddReview(r)
		_, _ = resp.Write([]byte("reviews are added."))
	}
}

func getReviews(rs review.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		reviews, _ := rs.GetReviews()
		reviewsBytes, _ := json.MarshalIndent(reviews, "", "  ")
		_, _ = resp.Write(reviewsBytes)
	}
}

func New(bs beer.Service, rs review.Service) *server {
	router := mux.NewRouter()
	router.HandleFunc("/beers", getBeers(bs)).Methods("GET")
	router.HandleFunc("/beers", addBeer(bs)).Methods("POST")
	router.HandleFunc("/reviews", getReviews(rs)).Methods("GET")
	router.HandleFunc("/reviews", addReview(rs)).Methods("POST")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return &server{
		server: &http.Server{
			Addr:    ":7777",
			Handler: router,
		},
		done: done,
	}
}
