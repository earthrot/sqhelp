package sqhelp

import (
	"testing"
	"time"
)

func TestNullTimeValue(t *testing.T) {
	var value time.Time = time.Now()

	res := NullTime(value)

	if res.Valid != true || res.Time != value {
		t.Fatalf("NullTime result not matching, got %s (%t) expected %s (%t)", res.Time, res.Valid, value, true)
	}
}

func TestNullTimeInvalid(t *testing.T) {
	res := NullTime(time.Time{}, false)

	if res.Valid != false {
		t.Fatalf("NullTime validity not matching, got %t expected %t", res.Valid, false)
	}
}
