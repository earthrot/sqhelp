package sqhelp

import (
	"math"
	"testing"
)

func TestNullFloat64Value(t *testing.T) {
	var value float64 = math.MaxFloat64

	res := NullFloat64(value)

	if res.Valid != true || res.Float64 != value {
		t.Fatalf("NullFloat64 result not matching, got %.2f (%t) expected %.2f (%t)", res.Float64, res.Valid, value, true)
	}
}

func TestNullFloat64Invalid(t *testing.T) {
	res := NullFloat64(0, false)

	if res.Valid != false {
		t.Fatalf("NullFloat64 validity not matching, got %t expected %t", res.Valid, false)
	}
}
