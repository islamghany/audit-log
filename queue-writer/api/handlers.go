package api

import (
	"context"
	"encoding/json"
	"net/http"
	"queue-writer/db/data"
	"queue-writer/internals/validator"
	"time"
)

var ctx = context.Background()

func (s *Server) createLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.notFoundResponse(w, r)
		return
	}
	var input struct {
		CreatedAt   time.Time `json:"created_at"`
		EventName   string    `json:"event_name"`
		Description string    `json:"description"`
		CustomerID  int64     `json:"customer_id"`
	}

	err := s.readJSON(w, r, &input)
	if err != nil {
		s.badRequestResponse(w, r, err)
		return
	}

	log := &data.Log{
		EventName:   input.EventName,
		Description: input.Description,
		CreatedAt:   input.CreatedAt,
		CustomerID:  input.CustomerID,
	}

	if log.CreatedAt.IsZero() {
		log.CreatedAt = time.Now()
	}

	v := validator.NewValidator()

	if data.ValidateLog(v, log); !v.Valid() {
		s.failedValidationResponse(w, r, v.Errors)
		return
	}

	payload, err := json.Marshal(log)
	if err != nil {
		s.serverErrorResponse(w, r, err)
		return
	}
	err = s.publisher.Publish(ctx, "audit-log", payload)
	if err != nil {
		s.serverErrorResponse(w, r, err)
		return
	}

	err = s.writeJSON(w, http.StatusCreated, envelope{"log": log}, nil)
	if err != nil {
		s.serverErrorResponse(w, r, err)
	}
}
