package api

import (
	"fmt"
	"net/http"
	"queue-writer/broker"
	"time"
)

type Server struct {
	publisher broker.Publisher
}

func NewServer(publisher broker.Publisher) *Server {
	return &Server{
		publisher: publisher,
	}
}

func (s *Server) Serve(port string) error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      s.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return srv.ListenAndServe()
}
