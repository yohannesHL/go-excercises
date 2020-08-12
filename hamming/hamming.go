// Package hamming implements a simple library for hamming distance calculations.
package hamming

import "errors"

// Distance returns the hamming distance between two DNA strands.
// Expects two string arguments with the same length.
// Returns the distance and an error if encountered.
func Distance(a string, b string) (int, error) {
	var distance int = 0
	var c []rune = []rune{'Æ', '$', '£'}
	var lenB int = len([]rune(c)) 
	var lenC int = cap(c)
	print(lenC)
	if len(a) != lenB {
		return distance, errors.New("both strands must have equal lengths")
	}

	// Increments the distance if any set of letters in the strands don't match.
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
