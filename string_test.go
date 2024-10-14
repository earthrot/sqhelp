package sqhelp

import "testing"

func TestNullStringValue(t *testing.T) {
	var value string = "foo"

	res := NullString(value)

	if res.Valid != true || res.String != value {
		t.Fatalf("NullString result not matching, got %s (%t) expected %s (%t)", res.String, res.Valid, value, true)
	}
}

func TestNullStringInvalid(t *testing.T) {
	res := NullString("", false)

	if res.Valid != false {
		t.Fatalf("NullString validity not matching, got %t expected %t", res.Valid, false)
	}
}
