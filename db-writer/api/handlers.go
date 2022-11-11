package api

import (
	"db-writer/db/data"
	"encoding/json"
	"fmt"
	"time"
)

func (s *Server) CreateLogHandler(payload string) {

	var input struct {
		CreatedAt   time.Time `json:"created_at"`
		EventName   string    `json:"event_name"`
		Description string    `json:"description"`
		CustomerID  int64     `json:"customer_id"`
	}

	err := json.Unmarshal([]byte(payload), &input)

	if err != nil {
		doSomethingWithAnInternalError(err)
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
	if err != nil {
		doSomethingWithAnInternalError(err)
		return
	}

	err = s.models.Log.Insert(log)
	if err != nil {
		doSomethingWithAnInternalError(err)
		return
	}
}

// inform that there is an error
// we can sent it via mail server for example
func doSomethingWithAnInternalError(err error) {
	fmt.Println(err)
}
