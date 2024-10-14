package sqhelp

import "database/sql"

// valid is optional and defaults to "true"
func Null(value any, valid ...bool) sql.Null[any] {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.Null[any]{
		Valid: isValid,
		V:     value,
	}
}
