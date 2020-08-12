// Package scrabble implements a simple library for the game scrabble.
package scrabble

import "unicode"

// Score returns the scrabble score of a given word.
func Score(word string) int {
	var score int
	for _, runeValue := range word {
		switch unicode.ToLower(runeValue) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}
	return score
}
