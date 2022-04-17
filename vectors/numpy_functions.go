package vectors

import (
	"errors"
	"math"
)

// Number is an interface for type float64 and int
type Number interface {
	~float64 | int
}

// ConvertFloat is a function that converts a Number type slice to a float64 slice
func ConvertFloat[T Number](slice []T) []float64 {
	var result []float64
	for _, v := range slice {
		result = append(result, float64(v))
	}
	return result
}

// Contains is a function that checks if a value is in a slice
func Contains[T Number](slice []T, val T) bool {
	for _, elem := range slice {
		if elem == val {
			return true
		}
	}
	return false
}

// Unique is a function that returns a slice of unique values
func Unique[T Number](data []T) []float64 {
	var uniqueData []T

	for _, value := range data {
		if !Contains(uniqueData, value) {
			uniqueData = append(uniqueData, value)
		}
	}

	return ConvertFloat(uniqueData)
}

// MultiplyBy is a function that multiplies a slice of numbers by a number
func MultiplyBy[T Number](array []T, factor T) []float64 {
	var multipliedArray []T
	for i := range array {
		multipliedArray = append(multipliedArray, array[i]*factor)
	}

	return ConvertFloat(multipliedArray)
}

// IncrementBy is a function that increments a slice of numbers by a number
func IncrementBy[T Number](array []T, factor T) []float64 {
	var incrementedArray []T
	for i := range array {
		incrementedArray = append(incrementedArray, array[i]+factor)
	}

	return ConvertFloat(incrementedArray)
}

// PowBy is a function that take power of a slice of numbers by a number
func PowBy[T Number](array []T, factor T) []float64 {
	var powArray []T
	for i := range array {
		powArray = append(powArray, T(math.Pow(float64(array[i]), float64(factor))))
	}

	return ConvertFloat(powArray)
}

// Apply is a function that applies a function to a slice of numbers
func Apply[T Number](array []T, f func(T) T) []float64 {
	var appliedArray []T
	for i := range array {
		appliedArray = append(appliedArray, f(array[i]))
	}

	return ConvertFloat(appliedArray)
}

// Sum is a function that sums a slice of numbers
func Sum[T Number](array []T) float64 {
	var sum T

	for _, value := range array {
		sum = sum + value
	}

	return float64(sum)
}

// Cumsum is a function that returns cumulative sums of a slice
func Cumsum[T Number](array []T) []float64 {
	var cumsum []T

	cumsum = append(cumsum, array[0])

	for i := 1; i < len(array); i++ {
		cumsum = append(cumsum, array[i]+cumsum[i-1])
	}

	return ConvertFloat(cumsum)
}

// Round is a function that rounds a slice of numbers to a given decimal
func Round[T Number](array []T, decimals int) []float64 {
	var roundedArray []T
	for i := range array {
		roundedNum := T(math.Round(float64(array[i])*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals)))
		roundedArray = append(roundedArray, roundedNum)
	}

	return ConvertFloat(roundedArray)
}

// Max is a function that returns the maximum value and its index of a slice of numbers
func Max[T Number](array []T) (float64, int) {
	max := array[0]
	index := 0
	for i, value := range array {
		if value > max {
			max = value
			index = i
		}
	}

	return float64(max), index
}

// Min is a function that returns the maximum value and its index of a slice of numbers
func Min[T Number](array []T) (float64, int) {
	min := array[0]
	index := 0

	for i, value := range array {
		if value < min {
			min = value
			index = i
		}
	}

	return float64(min), index
}

// Mean is a function that returns the mean of a slice of numbers
func Mean[T Number](array []T) float64 {
	return Sum(array) / float64(len(array))
}

// Abs is a function that returns the absolute value of a slice of numbers
func Abs[T Number](array []T) []float64 {
	var absArray []T
	for i := range array {
		absArray = append(absArray, T(math.Abs(float64(array[i]))))
	}

	return ConvertFloat(absArray)
}

// Arange is a function that returns a slice of numbers from start to end with a step
func Arange[T Number](start, stop, step T) []float64 {
	var array []T

	for i := start; i < stop; i += step {
		array = append(array, i)
	}

	return ConvertFloat(array)
}

// Where is a function that returns the indices and elements of a slice of numbers that satisfy a condition
func Where[T Number](array []T, f func(T) bool) ([]int, []float64) {
	var indices []int
	var elements []T

	for i, value := range array {
		if f(value) {
			indices = append(indices, i)
			elements = append(elements, value)
		}
	}

	return indices, ConvertFloat(elements)
}

// SumWith is a function that sums elements of two slices with each other
func SumWith[T Number](array1 []T, array2 []T) ([]float64, error) {
	var array []T
	var err error
	if len(array1) != len(array2) {
		err = errors.New("Arrays must be of the same length")
		return ConvertFloat(array), err
	}
	for i := range array1 {
		array = append(array, array1[i]+array2[i])
	}
	return ConvertFloat(array), err
}

// MultiplyWith is a function that multiplies elements of two slices with each other
func MultiplyWith[T Number](array1 []T, array2 []T) ([]float64, error) {
	var array []T
	var err error
	if len(array1) != len(array2) {
		err = errors.New("Arrays must be of the same length")
		return ConvertFloat(array), err
	}
	for i := range array1 {
		array = append(array, array1[i]*array2[i])
	}
	return ConvertFloat(array), err
}

// DividedBy is a function that divides elements of two slices with each other
func DividedBy[T Number](array1 []T, array2 []T) ([]float64, error) {
	var array []T
	var err error
	if len(array1) != len(array2) {
		err = errors.New("Arrays must be of the same length")
		return ConvertFloat(array), err
	}
	for i := range array1 {
		array = append(array, array1[i]/array2[i])
	}
	return ConvertFloat(array), err
}

// Zeros is a function that returns a slice of zeros of a given length
func Zeros[T Number](length T) []float64 {
	var array []T
	for i := 0; i < int(length); i++ {
		array = append(array, T(0))
	}
	return ConvertFloat(array)
}

// Ones is a function that returns a slice of ones of a given length
func Ones[T Number](length T) []float64 {
	var array []T
	for i := 0; i < int(length); i++ {
		array = append(array, T(1))
	}
	return ConvertFloat(array)
}

// Insert is a function that inserts an element into a slice at a given index
func Insert[T Number](array []T, index int, value T) []float64 {
	var newArray []T
	for i := 0; i < index; i++ {
		newArray = append(newArray, array[i])
	}
	newArray = append(newArray, value)
	for i := index; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return ConvertFloat(newArray)
}

// Delete is a function that deletes an element from a slice at a given index
func Delete[T Number](array []T, index int) []float64 {
	var newArray []T
	for i := 0; i < index; i++ {
		newArray = append(newArray, array[i])
	}
	for i := index + 1; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return ConvertFloat(newArray)
}
