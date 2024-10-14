package sqhelp

import (
	"testing"
)

func TestNullBoolValue(t *testing.T) {
	var value bool = true

	res := NullBool(value)

	if res.Valid != true || res.Bool != value {
		t.Fatalf("NullBool result not matching, got %t (%t) expected %t (%t)", res.Bool, res.Valid, value, true)
	}
}

func TestNullBoolInvalid(t *testing.T) {
	res := NullBool(false, false)

	if res.Valid != false {
		t.Fatalf("NullBool validity not matching, got %t expected %t", res.Valid, false)
	}
}
