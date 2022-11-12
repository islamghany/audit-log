package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Log LogModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Log: LogModel{DB: db},
	}
}
