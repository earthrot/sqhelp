package sqhelp

import (
	"database/sql"
	"time"
)

// valid is optional and defaults to "true"
func NullTime(value time.Time, valid ...bool) sql.NullTime {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullTime{
		Valid: isValid,
		Time:  value,
	}
}
