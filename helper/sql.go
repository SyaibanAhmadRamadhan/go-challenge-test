package helper

import (
	"database/sql"
)

func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func GetNullString(s sql.NullString) *string {
	if s.String == "" {
		return nil
	}

	return &s.String
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

func GetNullInt64(s sql.NullInt64) *int64 {
	if s.Int64 == 0 {
		return nil
	}

	return &s.Int64
}
