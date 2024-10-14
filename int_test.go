package sqhelp

import (
	"testing"
)

func TestNullInt16Value(t *testing.T) {
	var value int16 = int16(^uint16(0) >> 1)

	res := NullInt16(value)

	if res.Valid != true || res.Int16 != value {
		t.Fatalf("NullInt16 result not matching, got %d (%t) expected %d (%t)", res.Int16, res.Valid, value, true)
	}
}

func TestNullInt16Invalid(t *testing.T) {
	res := NullInt16(0, false)

	if res.Valid != false {
		t.Fatalf("NullInt64 validity not matching, got %t expected %t", res.Valid, false)
	}
}

func TestNullInt32Value(t *testing.T) {
	var value int32 = int32(^uint32(0) >> 1)

	res := NullInt32(value)

	if res.Valid != true || res.Int32 != value {
		t.Fatalf("NullInt32 result not matching, got %d (%t) expected %d (%t)", res.Int32, res.Valid, value, true)
	}
}

func TestNullInt32Invalid(t *testing.T) {
	res := NullInt32(0, false)

	if res.Valid != false {
		t.Fatalf("NullInt32 validity not matching, got %t expected %t", res.Valid, false)
	}
}

func TestNullInt64Value(t *testing.T) {
	var value int64 = int64(^uint64(0) >> 1)

	res := NullInt64(value)

	if res.Valid != true || res.Int64 != value {
		t.Fatalf("NullInt64 result not matching, got %d (%t) expected %d (%t)", res.Int64, res.Valid, value, true)
	}
}

func TestNullInt64Invalid(t *testing.T) {
	res := NullInt64(0, false)

	if res.Valid != false {
		t.Fatalf("NullInt64 validity not matching, got %t expected %t", res.Valid, false)
	}
}
