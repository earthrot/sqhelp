package sqhelp

import "database/sql"

// valid is optional and defaults to "true"
func NullByte(value byte, valid ...bool) sql.NullByte {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullByte{
		Byte:  value,
		Valid: isValid,
	}
}
