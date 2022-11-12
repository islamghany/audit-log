package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"querying-logs/db/data"
	"time"
)

type Server struct {
	db data.Models
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: data.NewModels(db),
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
