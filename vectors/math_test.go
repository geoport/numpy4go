package vectors

import (
	"math"
	"reflect"
	"testing"
)

func TestMultiplyBy(t *testing.T) {
	expected := []float64{2.2, 2.4, 2.6, 2.8, 3.0, 2.2}
	output := MultiplyBy(testSliceFloat, 2.0)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("MultiplyBy(%v,%d) = %v, want %v", testSliceFloat, 2, output, expected)
	}
}

func TestPow(t *testing.T) {
	expected := []float64{1, 4, 9, 16, 25, 9}
	output := Pow(ConvertFloat(testSliceInt), 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("PowBy(%v,%d) = %v, want %v", ConvertFloat(testSliceInt), 2, output, expected)
	}
}

func TestPowOf(t *testing.T) {
	expected := []float64{2, 4, 8}
	output := PowOf(2, []float64{1, 2, 3})
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("PowBy(%v,%d) = %v, want %v", ConvertFloat(testSliceInt), 2, output, expected)
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
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Round(%v) = %v, want %v", testSliceFloat, output, expected)
	}
}

func TestMax(t *testing.T) {
	expectedMax := 5.0
	expectedArgMax := 4
	outputMax, argMax := Max(ConvertFloat(testSliceInt))
	if outputMax != expectedMax && argMax != expectedArgMax {
		t.Errorf("Max(%v) = %v, want %v", ConvertFloat(testSliceInt), outputMax, expectedMax)
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
	expected := 3.0
	output := Mean(ConvertFloat(testSliceInt))
	if output != expected {
		t.Errorf("Mean(%v) = %v, want %v", ConvertFloat(testSliceInt), output, expected)
	}
}

func TestAbs(t *testing.T) {
	testSlice := MultiplyBy(testSliceFloat, -1.0)
	expected := testSliceFloat

	output := Abs(testSlice)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Abs(%v) = %v, want %v", testSlice, output, expected)
	}
}

func TestSumWith(t *testing.T) {
	expected := []float64{2.2, 2.4, 2.6, 2.8, 3, 2.2}
	output := SumWith(testSliceFloat, testSliceFloat)

	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("SumWith(%v,%v) = %v, want %v", testSliceFloat, testSliceFloat, output, expected)
	}
}

func TestDividedBy(t *testing.T) {
	expected := []float64{1, 1, 1, 1, 1, 1}
	output := DividedBy(testSliceFloat, testSliceFloat)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("DividedBy(%v,%v) = %v, want %v", testSliceFloat, testSliceFloat, output, expected)
	}
}

func TestGeomspace(t *testing.T) {
	expected1 := []float64{1, 10, 100, 1000}
	output1 := Round(Geomspace(1, 1000, 4), 1)
	if reflect.DeepEqual(expected1, output1) != true {
		t.Errorf("Geomspace(%v,%v,%d) = %v, want %v", 1, 1000, 4, output1, expected1)
	}

	expected2 := []float64{1, 2, 4, 8, 16}
	output2 := Round(Geomspace(1, 16, 5), 1)
	if reflect.DeepEqual(expected2, output2) != true {
		t.Errorf("Geomspace(%v,%v,%d) = %v, want %v", 1, 16, 5, output2, expected2)
	}
}

func TestDot(t *testing.T) {
	expected := 9.76
	output := Dot(testSliceFloat, testSliceFloat)
	if math.Round(expected) != math.Round(output) {
		t.Errorf("Dot(%v,%v) = %v, want %v", testSliceFloat, testSliceFloat, output, expected)
	}
}

func TestAngle(t *testing.T) {
	expected := []float64{0, 1.57, 0.79}
	output := Round(Angle(testSliceComplex), 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Angle(%v) = %v, want %v", testSliceComplex, output, expected)
	}
}

func TestReal(t *testing.T) {
	expected := []float64{1, 0, 1}
	output := Real(testSliceComplex)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Real(%v) = %v, want %v", testSliceComplex, output, expected)
	}
}

func TestImaginary(t *testing.T) {
	expected := []float64{0, 1, 1}
	output := Imaginary(testSliceComplex)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestConj(t *testing.T) {
	expected := []complex128{1 + 0i, 0 - 1i, 1 - 1i}
	output := Conj(testSliceComplex)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Conj(%v) = %v, want %v", testSliceComplex, output, expected)
	}
}

func TestLinSpace(t *testing.T) {
	expected := []float64{2., 2.25, 2.5, 2.75, 3}
	output := Round(LinSpace(2, 3, 5), 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestLogSpace(t *testing.T) {
	expected := []float64{100, 215.44, 464.16, 1000}
	output := Round(LogSpace(2, 3, 4), 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestAllClose(t *testing.T) {
	output := AllClose(testSliceFloat, testSliceFloat, 1e-10)
	if !output {
		t.Errorf("Got %v, want %v", output, true)
	}
}

func TestRealIfClose(t *testing.T) {
	expected := []float64{2.1, 5.2}
	testSlice := []complex128{2.1 + 4e-15i, 5.2 + 3.2e-16i}
	output := RealIfClose(testSlice)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestFloorDivide(t *testing.T) {
	expected := []float64{1, 1, 1, 0, 0, 0}
	testDenominator := []float64{1, 1.2, 1.3, 2.9, 3.1, 4.4}
	output := FloorDivide(testSliceFloat, testDenominator)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestDiff(t *testing.T) {
	expected := []float64{0.1, 0.1, 0.1, 0.1, -0.4}
	output := Round(Diff(testSliceFloat), 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestIsFinite(t *testing.T) {
	output := IsFinite(testSliceFloat)
	expected := []bool{true, true, true, true, true, true}
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestMatmul(t *testing.T) {
	array1 := [][]float64{
		{1, 0},
		{0, 1},
	}
	array2 := [][]float64{
		{4, 1},
		{2, 2},
	}
	expected := [][]float64{
		{4, 1},
		{2, 2},
	}
	output := Matmul(array1, array2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestInterp(t *testing.T) {
	xp := []float64{1, 2, 3}
	fp := []float64{3, 2, 0}
	x := []float64{0, 1, 1.5, 2.72, 3.14}
	expected := []float64{3, 3, 2.5, 0.56, 0}
	output := Interp(x, xp, fp)
	if reflect.DeepEqual(expected, Round(output, 2)) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestInverse(t *testing.T) {
	testMatrix := [][]float64{{1, 0}, {0, 2}}
	expected := [][]float64{
		{1, 0},
		{0, 0.5},
	}
	output := Inverse(testMatrix)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestPolyfit(t *testing.T) {
	expected := []float64{0.64, 0.55}
	x := []float64{1, 3, 4, 6, 8, 9, 11, 14}
	y := []float64{1, 2, 4, 4, 5, 7, 8, 9}
	a, b := Polyfit(x, y)
	output := []float64{a, b}
	if reflect.DeepEqual(expected, Round(output, 2)) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestMod(t *testing.T) {
	output := Mod(testSliceFloat, 0.5)
	expected := []float64{0.1, 0.2, 0.3, 0.4, 0., 0.1}
	if reflect.DeepEqual(expected, Round(output, 2)) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestUnwrap(t *testing.T) {
	testSlice := []float64{0, 0.78539816, 1.57079633, 5.49778714, 6.28318531}
	expected := []float64{0, 0.79, 1.57, -0.79, 0}
	output := Unwrap(testSlice)
	if reflect.DeepEqual(expected, Round(output, 2)) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}

func TestNorm(t *testing.T) {
	expected := []float64{3.12}
	output := Round([]float64{Norm(testSliceFloat)}, 2)
	if reflect.DeepEqual(expected, output) != true {
		t.Errorf("Got %v, want %v", output, expected)
	}
}
