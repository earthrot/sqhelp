package sqhelp

import "database/sql"

// valid is optional and defaults to "true"
func NullBool(value bool, valid ...bool) sql.NullBool {
	isValid := true
	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullBool{
		Valid: isValid,
		Bool:  value,
	}
}
