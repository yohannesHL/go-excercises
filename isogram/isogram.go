package isogram

import "unicode"

// IsIsogram returns true if a string has repeating letters false otherwise.
func IsIsogram(word string) bool {
	visited := make(map[rune]bool, len(word))

	for _, char := range word {
		key := unicode.ToLower(char)
		if !unicode.IsLetter(key) {
			continue
		}
		if visited[key] {
			return false
		}
		visited[key] = true
	}
	return true
}
