package data

import "database/sql"

type Models struct {
	Log LogModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Log: LogModel{DB: db},
	}
}

func NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}
