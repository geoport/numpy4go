package vectors

import (
	"reflect"
	"testing"
)

func TestCumtrapz(t *testing.T) {
	expected := []float64{1.5, 4, 7.5, 12}
	x := []float64{1, 2, 3, 4, 5}
	f := []float64{1, 2, 3, 4, 5}
	output := Cumtrapz(x, f)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}
