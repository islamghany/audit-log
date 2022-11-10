package data

import (
	"context"
	"database/sql"
	"db-writer/internals/validator"
	"time"
)

type Log struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	EventName   string    `json:"event_name"`
	Description string    `json:"description"`
	CustomerID  int64     `json:"customer_id,omitempty"`
}

func ValidateLog(v *validator.Validator, l *Log) {

	v.Check(l.EventName != "", "event_name", "must be provided")
	v.Check(len(l.EventName) <= 100, "event_name", "must be not more than 100 bytes long")

	v.Check(l.EventName != "", "description", "must be provided")
	v.Check(len(l.EventName) <= 1000, "description", "must be not more than 1000 bytes long")

	v.Check(l.CustomerID >= 0, "customer_id", "must be a valid id")

	v.Check(l.CreatedAt.Before(time.Now()), "created_at", "must be correct, not in the future")
}

type LogModel struct {
	DB *sql.DB
}

func (l *LogModel) Insert(log *Log) error {
	query := `
		INSERT INTO logs
		(event_name, description, created_at, customer_id)
		VALUES($1, $2, $3, $4)
		RETURNING id, created_at;
	`

	args := []interface{}{log.EventName, log.Description, log.CreatedAt, NewNullInt64(log.CustomerID)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return l.DB.QueryRowContext(ctx, query, args...).Scan(&log.ID, &log.CreatedAt)
}
