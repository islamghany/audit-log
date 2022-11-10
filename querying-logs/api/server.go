package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
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
