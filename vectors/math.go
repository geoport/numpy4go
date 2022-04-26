package vectors

import (
	"errors"
	"math"
)

// MultiplyBy multiplies elements of a slice with a number or slice of numbers
func MultiplyBy[T Float](array []float64, factor T) []float64 {
	var multipliedArray []float64
	if IsVector(factor) {
		factor := AsSlice(factor)
		for i := range array {
			multipliedArray = append(multipliedArray, array[i]*factor[i])
		}
	} else {
		for i := range array {
			multipliedArray = append(multipliedArray, array[i]*asFloat64(factor))
		}
	}

	return multipliedArray
}

// Pow take power of a slice of numbers by a number or slice of numbers
func Pow[T Float](base []float64, factor T) []float64 {
	var powArray []float64
	if IsVector(factor) {
		factor := AsSlice(factor)
		for i := range base {
			powArray = append(powArray, math.Pow(base[i], factor[i]))
		}
	} else {
		for i := range base {
			powArray = append(powArray, math.Pow(base[i], asFloat64(factor)))
		}
	}

	return powArray
}

// PowOf take power of a number or slice of numbers by slice of numbers
func PowOf[T Float](base T, factor []float64) []float64 {
	var powArray []float64
	if IsVector(base) {
		base := AsSlice(base)
		for i := range base {
			powArray = append(powArray, math.Pow(base[i], factor[i]))
		}
	} else {
		for i := range factor {
			powArray = append(powArray, math.Pow(asFloat64(base), factor[i]))
		}
	}

	return powArray
}

// Sum sums a slice of numbers
func Sum(array []float64) float64 {
	var sum float64

	for _, value := range array {
		sum = sum + value
	}

	return sum
}

// Cumsum returns cumulative sums of a slice
func Cumsum(array []float64) []float64 {
	var cumsum []float64

	cumsum = append(cumsum, array[0])

	for i := 1; i < len(array); i++ {
		cumsum = append(cumsum, array[i]+cumsum[i-1])
	}

	return cumsum
}

// Round rounds a slice of numbers to a given decimal
func Round(array []float64, decimals int) []float64 {
	var roundedArray []float64
	for i := range array {
		roundedNum := math.Round(array[i]*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
		roundedArray = append(roundedArray, roundedNum)
	}

	return roundedArray
}

// Max returns the maximum value and its index of a slice of numbers
func Max(array []float64) (float64, int) {
	max := array[0]
	index := 0
	for i, value := range array {
		if value > max {
			max = value
			index = i
		}
	}

	return max, index
}

// Min returns the maximum value and its index of a slice of numbers
func Min(array []float64) (float64, int) {
	min := array[0]
	index := 0

	for i, value := range array {
		if value < min {
			min = value
			index = i
		}
	}

	return min, index
}

// Mean returns the mean of a slice of numbers
func Mean(array []float64) float64 {
	if len(array) == 0 {
		return 0
	}
	return Sum(array) / float64(len(array))
}

// Abs returns the absolute value of a slice of numbers
func Abs(array []float64) []float64 {
	var absArray []float64
	for i := range array {
		absArray = append(absArray, math.Abs(array[i]))
	}

	return absArray
}

// SumWith sums elements of two slices with each other
func SumWith[T Float](array1 []float64, factor T) ([]float64, error) {
	var array []float64
	var err error
	if IsVector(factor) {
		array2 := AsSlice(factor)
		if len(array1) != len(array2) {
			err = errors.New("arrays must be of the same length")
			return array, err
		}
		for i := range array1 {
			array = append(array, array1[i]+array2[i])
		}
	} else {
		for i := range array1 {
			array = append(array, array1[i]+asFloat64(factor))
		}
	}
	return array, err
}

// DividedBy divides elements of two slices with each other
func DividedBy(array1 []float64, array2 []float64) ([]float64, error) {
	var array []float64
	var err error
	if len(array1) != len(array2) {
		err = errors.New("arrays must be of the same length")
		return array, err
	}
	for i := range array1 {
		array = append(array, array1[i]/array2[i])
	}
	return array, err
}

// Geomspace returns a slice of numbers spaced evenly on a geometric progression
func Geomspace(start float64, end float64, length float64) []float64 {
	array := []float64{start}
	step := math.Pow(end/start, 1/(length-1))

	for i := 1; i < int(length); i++ {
		array = append(array, array[i-1]*step)
	}
	return array
}

// Dot returns the dot product of two slices
func Dot(array1 []float64, array2 []float64) (float64, error) {
	var err error
	if len(array1) != len(array2) {
		err = errors.New("arrays must be of the same length")
		return 0, err
	}
	var sum float64
	for i := range array1 {
		sum += array1[i] * array2[i]
	}
	return sum, err
}

// Angle returns the angle of the complex argument
func Angle(array []complex128) []float64 {
	var angleArray []float64
	for _, elem := range array {
		if imag(elem) == 0 && real(elem) == 0 {
			angleArray = append(angleArray, 0)
		} else {
			angleArray = append(angleArray, math.Atan2(imag(elem), real(elem)))
		}
	}
	return angleArray
}

// Real returns the real part of a complex argument
func Real(array []complex128) []float64 {
	var realArray []float64
	for _, elem := range array {
		realArray = append(realArray, real(elem))
	}
	return realArray
}

// Imaginary returns the imaginary part of a complex argument
func Imaginary(array []complex128) []float64 {
	var imaginaryArray []float64
	for _, elem := range array {
		imaginaryArray = append(imaginaryArray, imag(elem))
	}
	return imaginaryArray
}

// Conj returns the complex conjugate of a complex argument
func Conj(array []complex128) []complex128 {
	var conjArray []complex128
	for _, elem := range array {
		conjArray = append(conjArray, complex(real(elem), -imag(elem)))
	}
	return conjArray
}

// LinSpace returns a slice of numbers spaced evenly on a linear progression
func LinSpace(start float64, end float64, length float64) []float64 {
	array := []float64{start}
	step := (end - start) / (length - 1)

	for i := 1; i < int(length); i++ {
		array = append(array, array[i-1]+step)
	}
	return array
}

// LogSpace returns a slice of numbers spaced evenly on a logarithmic progression
func LogSpace(start float64, end float64, length float64) []float64 {
	y := LinSpace(start, end, length)
	return PowOf(10, y)
}

// AllClose returns true if all elements of two slices are within a tolerance of each other
func AllClose(array1 []float64, array2 []float64, tol float64) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i := range array1 {
		if math.Abs(array1[i]-array2[i]) > tol {
			return false
		}
	}
	return true
}

// RealIfClose returns the real part of a complex argument if the imaginary part is within a tolerance of zero
func RealIfClose(array []complex128) any {
	var realArray []float64
	imaginery := Imaginary(array)
	tol := 2.22e-14

	filterFunc := func(x float64) bool {
		return math.Abs(x) < tol
	}

	allSmaller := All(imaginery, filterFunc)
	if allSmaller {
		for _, elem := range array {
			if math.Abs(real(elem)) > tol {
				realArray = append(realArray, real(elem))
			}
		}
		return realArray
	} else {
		return array
	}

}

// FloorDivide returns the floor of the quotient of two arrays
func FloorDivide[T Float](numerator []float64, denominator T) ([]float64, error) {
	var result []float64
	var err error
	if IsVector(denominator) {
		denominator := AsSlice(denominator)
		if len(numerator) != len(denominator) {
			err = errors.New("numerator and denominator must have the same length")
			return result, err
		}
		for i := range numerator {
			result = append(result, math.Floor(numerator[i]/denominator[i]))
		}
	} else {
		for i := range numerator {
			result = append(result, math.Floor(numerator[i]/asFloat64(denominator)))
		}
	}
	return result, err
}

// Diff returns the n-th differences of the given array.
func Diff(array []float64) []float64 {
	var result []float64
	for i := 1; i < len(array); i++ {
		result = append(result, array[i]-array[i-1])
	}
	return result
}

// IsFinite returns a bool array, where true if input element is finite.
func IsFinite(array []float64) []bool {
	var result []bool
	for _, elem := range array {
		if math.IsInf(elem, 0) {
			result = append(result, false)
		} else {
			result = append(result, true)
		}
	}
	return result
}

// Matmul returns the matrix product of two arrays
func Matmul(a, b [][]float64) ([][]float64, error) {
	var result [][]float64
	var err error
	transposeB := Transpose(b)
	if !CheckConsistency(a) || !CheckConsistency(transposeB) {
		err = errors.New("all rows or columns must have the same length")
		return result, err
	}
	if len(a[0]) != len(Transpose(b)[0]) {
		err = errors.New("")
		return result, err
	}
	for _, row := range a {
		var newRow []float64
		for _, col := range transposeB {
			mult, multErr := Dot(row, col)
			if err != nil {
				return result, multErr
			} else {
				newRow = append(newRow, mult)
			}
		}
		result = append(result, newRow)
	}
	return result, err
}

// Interp returns an array of linearly interpolated values.
func Interp(x, xp, fp []float64) ([]float64, error) {
	var result []float64
	var err error
	if len(xp) != len(fp) {
		err = errors.New("xp and fp must have the same length")
		return result, err
	}
	for _, xi := range x {
		if xi < xp[0] {
			result = append(result, fp[0])
		} else if xi > xp[len(xp)-1] {
			result = append(result, fp[len(fp)-1])
		} else {
			for i := 0; i < len(xp)-1; i++ {
				if xi >= xp[i] && xi <= xp[i+1] {
					result = append(result, fp[i]+(fp[i+1]-fp[i])/(xp[i+1]-xp[i])*(xi-xp[i]))
				}
			}
		}
	}
	return result, err
}

// Inverse returns the inverse of a 2X2 matrix
func Inverse(array [][]float64) ([][]float64, error) {
	var err error
	if len(array[0]) != 2 || len(array) != 2 {
		err = errors.New("array must be a 2X2 matrix")
		return [][]float64{}, err
	}
	coef := 1 / (array[0][0]*array[1][1] - array[0][1]*array[1][0])
	result := [][]float64{{array[1][1] * coef, -array[0][1] * coef}, {-array[1][0] * coef, array[0][0] * coef}}
	return result, err
}

// Polyfit returns the coefficients of a polynomial of degree 1 that fits the data
func Polyfit(x, y []float64) (float64, float64, error) {
	// y = ax +b
	var err error
	var a float64
	var b float64
	if len(x) != len(y) {
		err = errors.New("x and y must have the same length")
		return a, b, err
	}
	sumX := Sum(x)
	sumY := Sum(y)
	sumX2 := Sum(Pow(x, 2))
	XY := MultiplyBy(x, y)
	sumXY := Sum(XY)

	A := [][]float64{{sumX, float64(len(x))}, {sumX2, sumX}}
	B := [][]float64{{sumY}, {sumXY}}
	invA, _ := Inverse(A)
	coeff, _ := Matmul(invA, B)
	a = coeff[0][0]
	b = coeff[1][0]
	return a, b, err
}

// Mod returns the element-wise remainder of division
func Mod[T Float](x []float64, y T) ([]float64, error) {
	var result []float64
	var err error
	if IsVector(y) {
		y := AsSlice(y)
		if len(x) != len(y) {
			err = errors.New("x and y must have the same length")
			return result, err
		}
		for i := range x {
			result = append(result, math.Mod(x[i], y[i]))
		}
	} else {
		for i := range x {
			result = append(result, math.Mod(x[i], asFloat64(y)))
		}
	}

	return result, err
}

// Unwrap This unwraps a signal p by changing elements which have an absolute difference from their predecessor of
//more than max(discont, period/2) to their period-complementary values.
func Unwrap(array []float64) []float64 {
	result := []float64{array[0]}
	dd := Diff(array)
	discont := math.Pi
	period := 2 * math.Pi
	intervalHigh := period / 2
	intervalLow := -period / 2
	s, _ := SumWith(dd, -intervalLow)
	ddmod, _ := Mod(s, period)
	ddmod, _ = SumWith(ddmod, intervalLow)
	for i, d := range ddmod {
		if d == intervalLow && dd[i] > 0 {
			ddmod[i] = intervalHigh
		}
	}
	phCorrect, _ := SumWith(ddmod, MultiplyBy(dd, -1))
	for i := range phCorrect {
		if Abs(dd)[i] < discont {
			ddmod[i] = 0
		}
	}
	summed, _ := SumWith(array[1:], Cumsum(phCorrect))
	result = append(result, summed...)
	return result
}

// Norm returns the norm of an array
func Norm(x []float64) float64 {
	return math.Sqrt(Sum(Pow(x, 2)))
}
