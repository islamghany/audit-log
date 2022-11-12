package api

import (
	"errors"
	"net/http"
	"querying-logs/db/data"
	"querying-logs/internals/validator"
)

func (s *Server) getLogsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		s.notFoundResponse(w, r)
		return
	}

	var input struct {
		EventName  string
		CustomerID int
		data.Filters
	}

	v := validator.NewValidator()

	q := r.URL.Query()

	input.EventName = s.readString(q, "event_name", "")
	input.CustomerID = s.readInt(q, "customer_id", 0, v)
	input.Filters.Page = s.readInt(q, "page", 1, v)
	input.Filters.PageSize = s.readInt(q, "page_size", 20, v)
	input.Filters.Sort = s.readString(q, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "event_name", "created_at", "-event_name", "-id", "-created_at"}

	input.Filters.ValidateFilters(v)

	if !v.Valid() {
		s.failedValidationResponse(w, r, v.Errors)
		return
	}

	logs, metadata, err := s.db.Log.GetAll(input.EventName, input.CustomerID, input.Filters)
	if err != nil {
		s.serverErrorResponse(w, r, err)
		return
	}
	err = s.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "logs": logs}, nil)
	if err != nil {
		s.serverErrorResponse(w, r, err)
	}

}

func (s *Server) getLogHandler(w http.ResponseWriter, r *http.Request) {

	id, err := s.readIDParam(r)
	if err != nil {
		s.notFoundResponse(w, r)
		return
	}

	log, err := s.db.Log.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			s.notFoundResponse(w, r)
		default:
			s.serverErrorResponse(w, r, err)
		}
		return
	}

	err = s.writeJSON(w, http.StatusOK, envelope{"log": log}, nil)

	if err != nil {
		s.serverErrorResponse(w, r, err)
	}
}
