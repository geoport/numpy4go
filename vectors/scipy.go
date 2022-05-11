package vectors

import "sort"

// Cumtrapz cumulatively integrates f(x) using the composite trapezoidal rule.
func Cumtrapz(x, f []float64) []float64 {

	switch {
	case len(x) != len(f):
		panic("integrate: slice length mismatch")
	case len(x) < 2:
		panic("integrate: input data too small")
	case !sort.Float64sAreSorted(x):
		panic("integrate: input must be sorted")
	}
	integral := []float64{0}
	for i := 0; i < len(x)-1; i++ {
		n := 0.5 * (x[i+1] - x[i]) * (f[i+1] + f[i])
		integral = append(integral, integral[i]+n)
	}

	return integral[1:]
}
