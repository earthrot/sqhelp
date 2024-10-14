package sqhelp

import "testing"

func TestNullByteValue(t *testing.T) {
	var value byte = byte('F')

	res := NullByte(value)

	if res.Valid != true || res.Byte != value {
		t.Fatalf("NullByte result not matching, got %d (%t) expected %d (%t)", res.Byte, res.Valid, value, true)
	}
}

func TestNullByteInvalid(t *testing.T) {
	res := NullByte(0, false)

	if res.Valid != false {
		t.Fatalf("NullByte validity not matching, got %t expected %t", res.Valid, false)
	}
}
