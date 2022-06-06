package vectors

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func TabToSpace(input string) string {
	var result []string

	for _, i := range input {
		switch {
		// all these considered as space, including tab \t
		// '\t', '\n', '\v', '\f', '\r',' ', 0x85, 0xA0
		case unicode.IsSpace(i):
			if result[len(result)-1] != " " {
				result = append(result, " ") // replace tab with space
			}
		case !unicode.IsSpace(i):
			result = append(result, string(i))
		}
	}
	return strings.Join(result, "")
}

// Loadtxt returns a slice of float64s from a file.
func Loadtxt(filepath string, start int, unpack bool) [][]float64 {
	f, err := os.Open(filepath)
	result := [][]float64{{}}

	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		if i >= start {
			line := TabToSpace(scanner.Text())
			for colIndex, c := range strings.Split(line, " ") {
				s, _ := strconv.ParseFloat(c, 64)
				if unpack {
					if len(result) <= colIndex {
						result = append(result, []float64{})
					}
					result[colIndex] = append(result[colIndex], s)
				} else {
					result[0] = append(result[0], s)
				}
			}
		}
		i++
	}

	return result
}
