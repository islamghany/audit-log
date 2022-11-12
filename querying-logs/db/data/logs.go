package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Log struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	EventName   string    `json:"event_name"`
	Description string    `json:"description"`
	CustomerID  int64     `json:"customer_id,omitempty"`
}

type LogModel struct {
	DB *sql.DB
}

func (l LogModel) GetAll(event_name string, customer_id int, filters Filters) ([]*Log, Metadata, error) {
	cid := ""
	if customer_id != 0 {
		cid = fmt.Sprintf("AND customer_id=%d", customer_id)
	}
	query := fmt.Sprintf(`
	SELECT count(*) OVER(),id, created_at , event_name,description, customer_id
	FROM logs
	WHERE (to_tsvector('simple', event_name) @@ plainto_tsquery('simple', $1) OR $1 = '') %s
	ORDER BY %s %s, id ASC
	LIMIT $2 OFFSET $3`, cid, filters.sortColumn(), filters.sortDirection())

	var customerID sql.NullInt64

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{event_name, filters.limit(), filters.offset()}
	rows, err := l.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()
	totalRecords := 0
	logs := []*Log{}

	for rows.Next() {
		var log Log

		err := rows.Scan(
			&totalRecords,
			&log.ID,
			&log.CreatedAt,
			&log.EventName,
			&log.Description,
			&customerID,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		log.CustomerID = customerID.Int64
		// Add the Movie struct to the slice.
		logs = append(logs, &log)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return logs, metadata, nil
}

func (l LogModel) Get(id int64) (*Log, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at , event_name,description, customer_id
		FROM logs
		WHERE logs.id = $1;
	`

	var log Log

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var customerID sql.NullInt64
	err := l.DB.QueryRowContext(ctx, query, id).Scan(
		&log.ID,
		&log.CreatedAt,
		&log.EventName,
		&log.Description,
		&customerID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &log, nil
}
