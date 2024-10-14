package sqhelp

import "database/sql"

// valid is optional and defaults to "true"
func NullInt16(value int16, valid ...bool) sql.NullInt16 {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullInt16{
		Valid: isValid,
		Int16: value,
	}
}

// valid is optional and defaults to "true"
func NullInt32(value int32, valid ...bool) sql.NullInt32 {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullInt32{
		Valid: isValid,
		Int32: value,
	}
}

// valid is optional and defaults to "true"
func NullInt64(value int64, valid ...bool) sql.NullInt64 {
	isValid := true

	if len(valid) > 0 {
		isValid = valid[0]
	}

	return sql.NullInt64{
		Valid: isValid,
		Int64: value,
	}
}
