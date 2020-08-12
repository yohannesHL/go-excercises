// Package leap provides functions for determining leap years
package leap

import "math"

// IsLeapYear checks if the given year is a leap year
func IsLeapYear(year int) bool {
	// must be divisible by 4
	if math.Mod(float64(year), 4) != 0.0 {
		return false
	}
	// and may be divisible by 100 if also divisible by 400
	if math.Mod(float64(year), 100) == 0 && math.Mod(float64(year), 400) == 0 {
		return true
	}
	// otherwise not divisible by 100
	if math.Mod(float64(year), 100) != 0 {
		return true
	}
	return false
}
