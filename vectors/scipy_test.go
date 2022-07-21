package vectors

import (
	"reflect"
	"testing"
)

func TestCumtrapz(t *testing.T) {
	expected := []float64{0, 0.75, 2, 3.75, 6}
	dx := 0.5
	f := []float64{1, 2, 3, 4, 5}
	output := Cumtrapz(f, dx, 0)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}
