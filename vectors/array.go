package vectors

import (
	"reflect"
)

type Float interface {
	~float64 | []float64 | int64 | []int64 | int | []int | int32 | []int32
}

// IsVector checks if a variable is a slice
func IsVector[T Float](obj T) bool {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Slice:
		return true
	default:
		return false
	}
}

// AsSlice converts a variable to a slice
func AsSlice[T Float](obj T) []float64 {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Slice:
		return any(obj).([]float64)
	default:
		return []float64{any(obj).(float64)}
	}
}

// asFloat64 converts a variable to a float64
func asFloat64[T Float](obj T) float64 {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Slice:
		return any(obj).([]float64)[0]
	case reflect.Int:
		return float64(any(obj).(int))
	case reflect.Int64:
		return float64(any(obj).(int))
	case reflect.Int32:
		return float64(any(obj).(int))
	default:
		return any(obj).(float64)
	}
}

// ConvertFloat converts a int type slice to a float64 slice
func ConvertFloat(slice []int) []float64 {
	var result []float64
	for _, v := range slice {
		result = append(result, float64(v))
	}
	return result
}

// ConvertInt converts a float64 type slice to a int slice
func ConvertInt(slice []float64) []int {
	var result []int
	for _, v := range slice {
		result = append(result, int(v))
	}
	return result
}

// Contains checks if a value is in a slice
func Contains[T float64 | string | int](slice []T, val T) bool {
	for _, elem := range slice {
		if elem == val {
			return true
		}
	}
	return false
}

// Unique returns a slice of unique values
func Unique(data []float64) []float64 {
	var uniqueData []float64

	for _, value := range data {
		if !Contains(uniqueData, value) {
			uniqueData = append(uniqueData, value)
		}
	}

	return uniqueData
}

// MultiplyBy multiplies a slice of numbers by a number

// Apply applies a function to a slice of numbers
func Apply(array []float64, f func(float642 float64) float64) []float64 {
	var appliedArray []float64
	for i := range array {
		appliedArray = append(appliedArray, f(array[i]))
	}

	return appliedArray
}

// Arange returns a slice of numbers from start to end with a step
func Arange(start, stop, step float64) []float64 {
	var array []float64

	for i := start; i < stop; i += step {
		array = append(array, i)
	}

	return array
}

// Where returns the indices and elements of a slice of numbers that satisfy a condition
func Where(array []float64, f func(float642 float64) bool) ([]int, []float64) {
	var indices []int
	var elements []float64

	for i, value := range array {
		if f(value) {
			indices = append(indices, i)
			elements = append(elements, value)
		}
	}

	return indices, elements
}

// Zeros returns a slice of zeros of a given length
func Zeros(row int, col int) [][]float64 {
	var array [][]float64
	for r := 0; r < row; r++ {
		var row_ []float64
		for c := 0; c < col; c++ {
			row_ = append(row_, float64(0))
		}
		array = append(array, row_)
	}
	return array
}

// Ones returns a slice of ones of a given length
func Ones(length int) []float64 {
	var array []float64
	for i := 0; i < length; i++ {
		array = append(array, float64(1))
	}
	return array
}

// Insert inserts an element into a slice at a given index
func Insert(array []float64, index int, value float64) []float64 {
	var newArray []float64
	if index < 0 || index > len(array)-1 {
		panic("index is out of range")
	}

	for i := 0; i < index; i++ {
		newArray = append(newArray, array[i])
	}
	newArray = append(newArray, value)
	for i := index; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return newArray
}

// Delete deletes an element from a slice at a given index
func Delete(array []float64, index int) []float64 {
	var newArray []float64
	if index < 0 || index > len(array)-1 {
		panic("index is out of range")
	}

	for i := 0; i < index; i++ {
		newArray = append(newArray, array[i])
	}
	for i := index + 1; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return newArray
}

// Any returns true if any element of a slice satisfies a condition
func Any(array []float64, f func(float642 float64) bool) bool {
	for i := range array {
		if f(array[i]) {
			return true
		}
	}
	return false
}

// All returns true if all elements of a slice satisfy a condition
func All(array []float64, f func(float642 float64) bool) bool {
	for i := range array {
		if !f(array[i]) {
			return false
		}
	}
	return true
}

// Size returns the total number of elements in a slice
func Size(array [][]float64) int {
	size := 0
	for _, elem := range array {
		size += len(elem)
	}

	return size
}

// ColumnStack takes a sequence of 1-D arrays and stack them as columns to make a single 2-D array
func ColumnStack(arrays ...[]float64) [][]float64 {
	var result [][]float64

	for i := range arrays[0] {
		var row []float64
		for _, array := range arrays {
			row = append(row, array[i])
		}
		result = append(result, row)
	}

	return result
}

// Transpose returns the transpose of a 2-D array
func Transpose(array [][]float64) [][]float64 {
	var result [][]float64
	for i := range array[0] {
		var row []float64
		for _, elem := range array {
			row = append(row, elem[i])
		}
		result = append(result, row)
	}
	return result
}

// IsComplex returns a bool array, where true if input element is complex.
func IsComplex[T ~float64 | complex128](array []T) []bool {
	var result []bool
	for _, elem := range array {
		if reflect.TypeOf(elem).Kind() == reflect.Complex128 {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

// SearchSorted returns the indices of the sorted array that contain elements in the input array
func SearchSorted(array []float64, vals []float64, side string) []int {
	var indices []int
	for _, val := range vals {
		var indice int
		if val < array[0] {
			indice = 0
		} else if val > array[len(array)-1] {
			indice = len(array) - 1
		} else {
			filterFunc := func(x float64) bool {
				return x >= val
			}
			indexes, _ := Where(array, filterFunc)
			if side == "left" {
				indice = indexes[0]
			} else {
				indice = indexes[len(indexes)-1]
			}
		}
		indices = append(indices, indice)
	}
	return indices
}

// RowStack takes a sequence of 1-D arrays and stack them as rows to make a single 2-D array
func RowStack(arrays ...[]float64) [][]float64 {
	var result [][]float64
	for _, array := range arrays {
		result = append(result, array)
	}
	return result
}

// Roll returns an array with elements that roll beyond the last position are re-introduced at the first.
func Roll(array []float64, shift int) []float64 {
	var result []float64
	for i := len(array) - shift; i < len(array); i++ {
		result = append(result, array[i])
	}
	for i := 0; i < len(array)-shift; i++ {
		result = append(result, array[i])
	}

	return result
}

// CheckConsistency returns true if length of all the rows are the same
func CheckConsistency(array [][]float64) bool {
	for i := 1; i < len(array); i++ {
		if len(array[i]) != len(array[i-1]) {
			return false
		}
	}
	return true
}

// Flipud returns an array with the elements reversed.
func Flipud(array []float64) []float64 {
	var result []float64
	for i := len(array) - 1; i >= 0; i-- {
		result = append(result, array[i])
	}
	return result
}

// Tile returns an array with the elements repeated the number of times given by the input array.
func Tile(array []float64, reps int) [][]float64 {
	var result [][]float64
	for i := 0; i < reps; i++ {
		result = append(result, array)
	}
	return result
}

// Meshgrid returns two arrays with the coordinates of the points in a meshgrid.
func Meshgrid(x, y []float64) ([][]float64, [][]float64) {
	var xGrid, yGrid [][]float64
	for i := 0; i < len(y); i++ {
		xGrid = append(xGrid, x)
	}
	for i := 0; i < len(x); i++ {
		yGrid = append(yGrid, y)
	}
	return xGrid, Transpose(yGrid)
}

//GetColumn returns a column of a 2-D array
func GetColumn(array [][]float64, index int) []float64 {
	var result []float64
	for _, row := range array {
		result = append(result, row[index])
	}
	return result
}

// Repat returns an array with the elements repeated the number of times given by the input value.
func Repeat(elem float64, reps int) []float64 {
	var result []float64
	for i := 0; i < reps; i++ {
		result = append(result, elem)
	}
	return result
}
