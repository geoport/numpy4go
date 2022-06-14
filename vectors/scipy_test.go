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

func TestInterp1D(t *testing.T) {
	expected := []float64{1, 2, 3, 4, 5}
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{1, 2, 3, 4, 5}
	xi := []float64{1, 2, 3, 4, 5}
	output := Interp1D(x, y, xi)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}
