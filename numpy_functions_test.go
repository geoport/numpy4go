package GeotechnicalSubroutines

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

var testSliceFloat = []float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.1}
var testSliceInt = []int{1, 2, 3, 4, 5, 3}
var testSliceAngle = []float64{math.Pi, math.Pi / 2}

func TestContains(t *testing.T) {
	testElemTrueInt := 1
	testElemFalseInt := 10
	testElemTrueFloat := 1.1
	testElemFalseFloat := 1.11

	if Contains(testSliceInt, testElemTrueInt) != true {
		t.Errorf("Contains(%v, %v) = %v, want %v", testSliceInt, testElemTrueInt, Contains(testSliceInt, testElemTrueInt), true)
	}
	if Contains(testSliceInt, testElemFalseInt) == true {
		t.Errorf("Contains(%v, %v) = %v, want %v", testSliceInt, testElemFalseInt, Contains(testSliceInt, testElemFalseInt), true)
	}

	if Contains(testSliceFloat, testElemTrueFloat) != true {
		t.Errorf("Contains(%v, %v) = %v, want %v", testSliceFloat, testElemTrueFloat, Contains(testSliceFloat, testElemTrueFloat), true)
	}
	if Contains(testSliceFloat, testElemFalseFloat) == true {
		t.Errorf("Contains(%v, %v) = %v, want %v", testSliceFloat, testElemFalseFloat, Contains(testSliceFloat, testElemFalseFloat), true)
	}
}

func TestUnique(t *testing.T) {
	expectedSlice := []int{1, 2, 3, 4, 5}
	outputSlice := Unique(testSliceInt)

	if reflect.DeepEqual(expectedSlice, outputSlice) != true {
		t.Errorf("Unique(%v) = %v, want %v", testSliceInt, outputSlice, expectedSlice)
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

func TestMultiplyBy(t *testing.T) {
	expected := []float64{2.2, 2.4, 2.6, 2.8, 3.0, 2.2}
	output := MultiplyBy(testSliceFloat, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("MultiplyBy(%v,%d) = %v, want %v", testSliceFloat, 2, output, expected)
	}
}

func TestIncrementBy(t *testing.T) {
	expected := []float64{3.1, 3.2, 3.3, 3.4, 3.5, 3.1}
	output := IncrementBy(testSliceFloat, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("IncrementBy(%v,%d) = %v, want %v", testSliceFloat, 2, output, expected)
	}
}

func TestPowBy(t *testing.T) {
	expected := []int{1, 4, 9, 16, 25, 9}
	output := PowBy(testSliceInt, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("PowBy(%v,%d) = %v, want %v", testSliceInt, 2, output, expected)
	}
}

func TestSum(t *testing.T) {
	expected := 7.6
	output := Sum(testSliceFloat)
	if output != expected {
		t.Errorf("Sum(%v) = %v, want %v", testSliceFloat, output, expected)
	}
}

func TestCumsum(t *testing.T) {
	expected := []float64{1.1, 2.3, 3.6, 5, 6.5, 7.6}
	output := Round(Cumsum(testSliceFloat), 1)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Cumsum(%v) = %v, want %v", testSliceFloat, output, expected)
	}
}

func TestRound(t *testing.T) {
	expected := []float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.1}
	output := Round(testSliceFloat, 1)
	fmt.Println(output)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Round(%v) = %v, want %v", testSliceFloat, output, expected)
	}
}

func TestMax(t *testing.T) {
	expectedMax := 5
	expectedArgMax := 4
	outputMax, argMax := Max(testSliceInt)
	if outputMax != expectedMax && argMax != expectedArgMax {
		t.Errorf("Max(%v) = %v, want %v", testSliceInt, outputMax, expectedMax)
	}
}

func TestMin(t *testing.T) {
	expectedMin := 1.1
	expectedArgMin := 0
	outputMin, argMin := Min(testSliceFloat)
	if outputMin != expectedMin && argMin != expectedArgMin {
		t.Errorf("Min(%v) = %v, want %v", testSliceFloat, outputMin, expectedMin)
	}
}

func TestMean(t *testing.T) {
	expected := 3
	output := Mean(testSliceInt)
	if output != expected {
		t.Errorf("Mean(%v) = %v, want %v", testSliceInt, output, expected)
	}
}

func TestAbs(t *testing.T) {
	testSlice := MultiplyBy(testSliceFloat, -1)
	expected := testSliceFloat

	output := Abs(testSlice)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Abs(%v) = %v, want %v", testSlice, output, expected)
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

func TestSumWith(t *testing.T) {
	expected := []float64{2.2, 2.4, 2.6, 2.8, 3, 2.2}
	output, _ := SumWith(testSliceFloat, testSliceFloat)

	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("SumWith(%v,%v) = %v, want %v", testSliceFloat, testSliceFloat, output, expected)
	}
}

func TestMultiplyWith(t *testing.T) {
	expected := PowBy(testSliceFloat, 2)
	output, _ := MultiplyWith(testSliceFloat, testSliceFloat)

	if reflect.DeepEqual(Round(expected, 2), Round(output, 2)) != true {
		t.Errorf("MultiplyWith(%v,%v) = %v, want %v", testSliceFloat, testSliceFloat, output, expected)
	}
}

func TestDividedBy(t *testing.T) {
	expected := []float64{1, 1, 1, 1, 1, 1}
	output, _ := DividedBy(testSliceFloat, testSliceFloat)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("DividedBy(%v,%v) = %v, want %v", testSliceFloat, testSliceFloat, output, expected)
	}
}

func TestOnes(t *testing.T) {
	expected := []int{1, 1, 1, 1, 1, 1}
	output := Ones(6)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Ones(%d) = %v, want %v", 6, output, expected)
	}
}

func TestZeros(t *testing.T) {
	expected := []int{0, 0, 0, 0, 0, 0}
	output := Zeros(6)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Zeros(%d) = %v, want %v", 6, output, expected)
	}
}

func TestInsert(t *testing.T) {
	expected := []int{1, 2, 2, 3, 4, 5, 3}
	output := Insert(testSliceInt, 1, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Insert(%v,%d,%d) = %v, want %v", testSliceInt, 1, 2, output, expected)
	}
}

func TestDelete(t *testing.T) {
	expected := []int{1, 3, 4, 5, 3}
	output := Delete(testSliceInt, 1)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Delete(%v,%d) = %v, want %v", testSliceInt, 1, output, expected)
	}
}
