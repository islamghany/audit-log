package api

import (
	"db-writer/db/data"
	"db-writer/internals/validator"
	"net/http"
	"time"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 3 // 3MB

func (s *Server) createLogHandler(w http.ResponseWriter, r *http.Request) {
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

	err = s.models.Log.Insert(log)
	if err != nil {
		s.serverErrorResponse(w, r, err)
		return
	}
	err = s.writeJSON(w, http.StatusCreated, envelope{"log": log}, nil)
	if err != nil {
		s.serverErrorResponse(w, r, err)
	}
}
