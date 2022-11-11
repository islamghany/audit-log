package data

import (
	"queue-writer/internals/validator"
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

	v.Check(l.CreatedAt.IsZero() || l.CreatedAt.Before(time.Now()), "created_at", "must be correct, not in the future")
}
