package sqhelp

import "database/sql"

// valid is optional and defaults to "true"
func NullFloat64(value float64, valid ...bool) sql.NullFloat64 {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullFloat64{
		Valid:   isValid,
		Float64: value,
	}
}
