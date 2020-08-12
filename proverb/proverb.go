// Package proverb is a library for generating proverbs
package proverb

import "fmt"

// Proverb creates a proverb from an array of words
func Proverb(rhyme []string) []string {
	var wordCount int = len(rhyme)
	var proverb []string

	if wordCount == 0 {
		return []string{}
	}

	for i := 1; i < wordCount; i++ {
		piece := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i])

		proverb = append(proverb, piece)
	}
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
