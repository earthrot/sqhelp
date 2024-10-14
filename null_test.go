package sqhelp

import "testing"

func TestNullValue(t *testing.T) {
	var value string = "bar"

	res := Null(value)

	if res.Valid != true || res.V != value {
		t.Fatalf("Null result not matching, got %s (%t) expected %s (%t)", res.V, res.Valid, value, true)
	}
}

func TestNullInvalid(t *testing.T) {
	res := Null(0, false)

	if res.Valid != false {
		t.Fatalf("Null validity not matching, got %t expected %t", res.Valid, false)
	}
}
