package http

import (
	"context"
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
		_ = bs.AddBeer(
			beer.Beer{
				Name:    "",
				Brewery: "",
				Abv:     0,
				Desc:    "",
			},
		)
		_, _ = resp.Write([]byte("beers are added."))
	}
}

func addReview(rs review.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		_ = rs.AddReview(review.Review{
			BeerID:       "",
			ReviewerName: "",
			Score:        0,
			Text:         "",
		})
		_, _ = resp.Write([]byte("reviews are added."))
	}
}

func New(bs beer.Service, rs review.Service) *server {
	router := mux.NewRouter()
	router.HandleFunc("/beers", addBeer(bs)).Methods("POST")
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
