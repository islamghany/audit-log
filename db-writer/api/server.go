package api

import (
	"database/sql"
	"db-writer/db/data"
)

type Server struct {
	models data.Models
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		models: data.NewModels(db),
	}
}
