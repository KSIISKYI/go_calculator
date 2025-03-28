package helpers

import "math"

// Round a float to a specific number of decimal places
// fix the precision issue
func Round(value float64, places int) float64 {
	shift := math.Pow(10, float64(places))

	return math.Round(value*shift) / shift
}
