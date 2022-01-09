package utils

import (
	"database/sql"
	"time"
)

func ConvSqlNullTime(sqlT sql.NullTime) time.Time {
	if sqlT.Valid {
		return sqlT.Time
	}

	return time.Time{}
}

func ConvTimeToSqlNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
