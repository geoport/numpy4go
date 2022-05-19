package vectors

import (
	"reflect"
	"testing"
)

func TestLoadtxt(t *testing.T) {
	expected1 := []float64{1, 2, 3, 4, 5, 6, 7}
	output1 := Loadtxt("testdata/single_column.txt", 0, false)[0]
	if reflect.DeepEqual(expected1, output1) != true {
		t.Errorf("Got %v, want %v", output1, expected1)
	}

	expected2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output2 := Loadtxt("testdata/multiple_column.txt", 0, false)[0]
	if reflect.DeepEqual(expected2, output2) != true {
		t.Errorf("Got %v, want %v", output2, expected2)
	}

	expected3 := [][]float64{
		{1, 4, 7, 10}, {2, 5, 8}, {3, 6, 9},
	}
	output3 := Loadtxt("testdata/multiple_column.txt", 0, true)
	if reflect.DeepEqual(expected3, output3) != true {
		t.Errorf("Got %v, want %v", output3, expected3)
	}
}
