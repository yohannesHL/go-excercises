// Package triangle is a simple library for dealing with triangles
package triangle

import "math"

// Kind represents type of triangle Equ = 3, this is -1 if not a triangle (NaT)
type Kind int

const (
	// NaT not a triangle
	NaT Kind = -1
	// Equ equilateral - all sides equal
	Equ Kind = 3
	// Iso isosceles - 2 sides equal
	Iso Kind = 2
	// Sca scalene - no sides equal
	Sca Kind = 0
)

// KindFromSides determines if the shape with 3 side lengths ( a, b, c )
// is not triangle (NaT), an scalene (Sca), isosceles (Iso) 
// or an equilateral (Equ) triangle. 
func KindFromSides(a, b, c float64) Kind {
	if isInValidRealNumber(a) || isInValidRealNumber(b) || isInValidRealNumber(c) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	// triangle inequality test
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

// isInValidRealNumber checks if number is not a valid real number
func isInValidRealNumber(number float64) bool {
	return math.IsNaN(number) || math.IsInf(number, 0)
}
