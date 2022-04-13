package GeotechnicalSubroutines

import (
	"errors"
	"math"
)

// Number is an interface for type float64 and int
type Number interface {
	~float64 | int
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
func Unique[T Number](data []T) []T {
	var uniqueData []T

	for _, value := range data {
		if !Contains(uniqueData, value) {
			uniqueData = append(uniqueData, value)
		}
	}

	return uniqueData
}

// MultiplyBy is a function that multiplies a slice of numbers by a number
func MultiplyBy[T Number](array []T, factor T) []T {
	var multipliedArray []T
	for i := range array {
		multipliedArray = append(multipliedArray, array[i]*factor)
	}

	return multipliedArray
}

// IncrementBy is a function that increments a slice of numbers by a number
func IncrementBy[T Number](array []T, factor T) []T {
	var incrementedArray []T
	for i := range array {
		incrementedArray = append(incrementedArray, array[i]+factor)
	}

	return incrementedArray
}

// PowBy is a function that take power of a slice of numbers by a number
func PowBy[T Number](array []T, factor T) []T {
	var powArray []T
	for i := range array {
		powArray = append(powArray, T(math.Pow(float64(array[i]), float64(factor))))
	}

	return powArray
}

// Apply is a function that applies a function to a slice of numbers
func Apply[T Number](array []T, f func(T) T) []T {
	var appliedArray []T
	for i := range array {
		appliedArray = append(appliedArray, f(array[i]))
	}

	return appliedArray
}

// Sum is a function that sums a slice of numbers
func Sum[T Number](array []T) T {
	var sum T

	for _, value := range array {
		sum = sum + value
	}

	return sum
}

// Cumsum is a function that returns cumulative sums of a slice
func Cumsum[T Number](array []T) []T {
	var cumsum []T

	cumsum = append(cumsum, array[0])

	for i := 1; i < len(array); i++ {
		cumsum = append(cumsum, array[i]+cumsum[i-1])
	}

	return cumsum
}

// Round is a function that rounds a slice of numbers to a given decimal
func Round[T Number](array []T, decimals int) []T {
	var roundedArray []T
	for i := range array {
		roundedNum := T(math.Round(float64(array[i])*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals)))
		roundedArray = append(roundedArray, roundedNum)
	}

	return roundedArray
}

// Max is a function that returns the maximum value and its index of a slice of numbers
func Max[T Number](array []T) (T, int) {
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

// Min is a function that returns the maximum value and its index of a slice of numbers
func Min[T Number](array []T) (T, int) {
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

// Mean is a function that returns the mean of a slice of numbers
func Mean[T Number](array []T) T {
	return Sum(array) / T(len(array))
}

// Abs is a function that returns the absolute value of a slice of numbers
func Abs[T Number](array []T) []T {
	var absArray []T
	for i := range array {
		absArray = append(absArray, T(math.Abs(float64(array[i]))))
	}

	return absArray
}

// Arange is a function that returns a slice of numbers from start to end with a step
func Arange[T Number](start, stop, step T) []T {
	var array []T

	for i := start; i < stop; i += step {
		array = append(array, i)
	}

	return array
}

// Where is a function that returns the indices and elements of a slice of numbers that satisfy a condition
func Where[T Number](array []T, f func(T) bool) ([]int, []T) {
	var indices []int
	var elements []T

	for i, value := range array {
		if f(value) {
			indices = append(indices, i)
			elements = append(elements, value)
		}
	}

	return indices, elements
}

// SumWith is a function that sums elements of two slices with each other
func SumWith[T Number](array1 []T, array2 []T) ([]T, error) {
	var array []T
	var err error
	if len(array1) != len(array2) {
		err = errors.New("Arrays must be of the same length")
		return array, err
	}
	for i := range array1 {
		array = append(array, array1[i]+array2[i])
	}
	return array, err
}

// MultiplyWith is a function that multiplies elements of two slices with each other
func MultiplyWith[T Number](array1 []T, array2 []T) ([]T, error) {
	var array []T
	var err error
	if len(array1) != len(array2) {
		err = errors.New("Arrays must be of the same length")
		return array, err
	}
	for i := range array1 {
		array = append(array, array1[i]*array2[i])
	}
	return array, err
}

// DividedBy is a function that divides elements of two slices with each other
func DividedBy[T Number](array1 []T, array2 []T) ([]T, error) {
	var array []T
	var err error
	if len(array1) != len(array2) {
		err = errors.New("Arrays must be of the same length")
		return array, err
	}
	for i := range array1 {
		array = append(array, array1[i]/array2[i])
	}
	return array, err
}

// Zeros is a function that returns a slice of zeros of a given length
func Zeros[T Number](length T) []T {
	var array []T
	for i := 0; i < int(length); i++ {
		array = append(array, T(0))
	}
	return array
}

// Ones is a function that returns a slice of ones of a given length
func Ones[T Number](length T) []T {
	var array []T
	for i := 0; i < int(length); i++ {
		array = append(array, T(1))
	}
	return array
}

// Insert is a function that inserts an element into a slice at a given index
func Insert[T Number](array []T, index int, value T) []T {
	var newArray []T
	for i := 0; i < index; i++ {
		newArray = append(newArray, array[i])
	}
	newArray = append(newArray, value)
	for i := index; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return newArray
}

// Delete is a function that deletes an element from a slice at a given index
func Delete[T Number](array []T, index int) []T {
	var newArray []T
	for i := 0; i < index; i++ {
		newArray = append(newArray, array[i])
	}
	for i := index + 1; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return newArray
}
