package api

import (
	"database/sql"
	"db-writer/db/data"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	models data.Models
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		models: data.NewModels(db),
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
