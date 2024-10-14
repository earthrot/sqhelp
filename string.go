package sqhelp

import "database/sql"

// valid is optional and defaults to "true"
func NullString(value string, valid ...bool) sql.NullString {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullString{
		Valid:  isValid,
		String: value,
	}
}
