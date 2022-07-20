package vectors

import (
	"math"
	"reflect"
	"testing"
)

var testSliceFloat = []float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.1}
var testSlice2D = [][]float64{{1.1, 1.2, 1.3, 1.4, 1.5, 1.1}, {1.1, 1.2, 1.3, 1.4, 1.5, 1.1}}
var testSliceInt = []int{1, 2, 3, 4, 5, 3}
var testSliceAngle = []float64{math.Pi, math.Pi / 2}
var testSliceComplex = []complex128{1, 1i, 1 + 1i}

func TestContains(t *testing.T) {
	testElemTrueFloat := 1.1
	testElemFalseFloat := 1.11

	if Contains(testSliceFloat, testElemTrueFloat) != true {
		t.Errorf("Contains(%v, %v) = %v, want %v", testSliceFloat, testElemTrueFloat, Contains(testSliceFloat, testElemTrueFloat), true)
	}
	if Contains(testSliceFloat, testElemFalseFloat) == true {
		t.Errorf("Contains(%v, %v) = %v, want %v", testSliceFloat, testElemFalseFloat, Contains(testSliceFloat, testElemFalseFloat), true)
	}
}

func TestUnique(t *testing.T) {
	expectedSlice := []float64{1.1, 1.2, 1.3, 1.4, 1.5}
	outputSlice := Unique(testSliceFloat)

	if reflect.DeepEqual(expectedSlice, outputSlice) != true {
		t.Errorf("Unique(%v) = %v, want %v", testSliceFloat, outputSlice, expectedSlice)
	}
}

func TestApply(t *testing.T) {
	f := math.Sin
	expected := []float64{math.Sin(math.Pi), math.Sin(math.Pi / 2)}
	output := Apply(testSliceAngle, f)

	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Apply(%v,%s) = %v, want %v", testSliceAngle, "sin", output, expected)
	}
}

func TestArange(t *testing.T) {
	expected := []float64{0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}
	output := Round(Arange(0.0, 1.0, 0.1), 1)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Arange(%d,%d,%f) = %v, want %v", 0, 1, 0.1, output, expected)
	}
}

func TestWhere(t *testing.T) {
	filterFunc := func(x float64) bool {
		return x >= 1.2 && x <= 1.4
	}
	expectedElems := []float64{1.2, 1.3, 1.4}
	expectedIndices := []int{1, 2, 3}
	outputInd, outputElems := Where(testSliceFloat, filterFunc)

	if reflect.DeepEqual(expectedElems, outputElems) != true && reflect.DeepEqual(expectedIndices, outputInd) != true {
		t.Errorf("Expected Indexes = %v, want %v. Expected Elems = %v, want %v", expectedIndices, outputInd, expectedElems, outputElems)
	}
}

func TestOnes(t *testing.T) {
	expected := []float64{1, 1, 1, 1, 1, 1}
	output := Ones(6)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Ones(%d) = %v, want %v", 6, output, expected)
	}
}

func TestZeros(t *testing.T) {
	expected := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	output := Zeros(3, 3)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Zeros(%d) = %v, want %v", 6, output, expected)
	}
}

func TestInsert(t *testing.T) {
	expected := []float64{1, 2, 2, 3, 4, 5, 3}
	output := Insert(ConvertFloat(testSliceInt), 1, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Insert(%v,%d,%d) = %v, want %v", ConvertFloat(testSliceInt), 1, 2, output, expected)
	}
}

func TestDelete(t *testing.T) {
	expected := []float64{1, 3, 4, 5, 3}
	output := Delete(ConvertFloat(testSliceInt), 1)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Delete(%v,%d) = %v, want %v", ConvertFloat(testSliceInt), 1, output, expected)
	}
}

func TestAny(t *testing.T) {
	filterFunc1 := func(x float64) bool {
		return x >= 1.2 && x <= 1.4
	}

	output1 := Any(testSliceFloat, filterFunc1)
	if output1 != true {
		t.Errorf("Got %v, want %v", output1, true)
	}

	filterFunc2 := func(x float64) bool {
		return x == 12
	}

	output2 := Any(testSliceFloat, filterFunc2)
	if output2 != false {
		t.Errorf("Got %v, want %v", output2, false)
	}
}

func TestAll(t *testing.T) {
	filterFunc1 := func(x float64) bool {
		return x >= 0
	}
	output1 := All(testSliceFloat, filterFunc1)
	if output1 != true {
		t.Errorf("Got %v, want %v", output1, true)
	}

	filterFunc2 := func(x float64) bool {
		return x == 12
	}

	output2 := All(testSliceFloat, filterFunc2)
	if output2 != false {
		t.Errorf("Got %v, want %v", output2, false)
	}
}

func TestSize(t *testing.T) {
	expected := 12
	testElem := [][]float64{testSliceFloat, testSliceFloat}
	output := Size(testElem)
	if expected != output {
		t.Errorf("Expected %v, want %v", output, expected)
	}
}

func TestColumnStack(t *testing.T) {
	expected := [][]float64{
		{1.1, 1.1},
		{1.2, 1.2},
		{1.3, 1.3},
		{1.4, 1.4},
		{1.5, 1.5},
		{1.1, 1.1},
	}
	output := ColumnStack(testSliceFloat, testSliceFloat)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestTranspose(t *testing.T) {
	expected := [][]float64{
		{1.1, 1.2, 1.3, 1.4, 1.5, 1.1},
		{1.1, 1.2, 1.3, 1.4, 1.5, 1.1},
	}
	output := Transpose(ColumnStack(testSliceFloat, testSliceFloat))
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestIsComplex(t *testing.T) {
	output := IsComplex(testSliceComplex)
	expected := []bool{true, true, true}
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestSearchSorted(t *testing.T) {
	expectedLeft := []int{0}
	expectedRight := []int{5}

	outputLeft := SearchSorted(testSliceFloat, []float64{1.1}, "left")
	outputRight := SearchSorted(testSliceFloat, []float64{1.1}, "right")

	if reflect.DeepEqual(expectedLeft, outputLeft) != true {
		t.Errorf("Got %v, want %v", outputLeft, expectedLeft)
	}
	if reflect.DeepEqual(expectedRight, outputRight) != true {
		t.Errorf("Got %v, want %v", outputRight, expectedRight)
	}
}

func TestRowStack(t *testing.T) {
	expected := [][]float64{
		{1.1, 1.2, 1.3, 1.4, 1.5, 1.1},
		{1.1, 1.2, 1.3, 1.4, 1.5, 1.1},
	}
	output := RowStack(testSliceFloat, testSliceFloat)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestRoll(t *testing.T) {
	expected := []float64{1.5, 1.1, 1.1, 1.2, 1.3, 1.4}
	output := Roll(testSliceFloat, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestFlipud(t *testing.T) {
	expected := []float64{1.1, 1.5, 1.4, 1.3, 1.2, 1.1}
	output := Flipud(testSliceFloat)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestTile(t *testing.T) {
	expected := [][]float64{testSliceFloat, testSliceFloat}
	output := Tile(testSliceFloat, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestCheckConsistency(t *testing.T) {
	testSliceFalse := [][]float64{testSliceFloat, testSliceFloat[1:]}
	testSliceTrue := [][]float64{testSliceFloat, testSliceFloat}

	if CheckConsistency(testSliceTrue) == false {
		t.Errorf("Got false, want true")
	}
	if CheckConsistency(testSliceFalse) {
		t.Errorf("Got true, want false")
	}
}

func TestMeshgrid(t *testing.T) {
	x := []float64{0, 1, 2}
	y := []float64{3, 4, 5}
	expectedX := [][]float64{
		{0, 1, 2},
		{0, 1, 2},
		{0, 1, 2},
	}
	expectedY := [][]float64{
		{3, 3, 3},
		{4, 4, 4},
		{5, 5, 5},
	}
	xGrid, yGrid := Meshgrid(x, y)
	if reflect.DeepEqual(expectedX, xGrid) != true {
		t.Errorf("Got %v, want %v", xGrid, expectedX)
	}
	if reflect.DeepEqual(expectedY, yGrid) != true {
		t.Errorf("Got %v, want %v", yGrid, expectedY)
	}
}

func TestGetColumn(t *testing.T) {
	testInput := [][]float64{testSliceFloat, testSliceFloat}
	expected := []float64{1.1, 1.1}
	output := GetColumn(testInput, 0)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestConvertInt(t *testing.T) {
	testInput := []float64{1, 2, 3}
	expected := []int{1, 2, 3}
	output := ConvertInt(testInput)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestRepeat(t *testing.T) {
	expected := []float64{2, 2, 2}
	output := Repeat(2, 3)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}
