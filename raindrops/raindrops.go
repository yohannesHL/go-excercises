package raindrops

import "strconv"

/*
Convert returns a string that contains raindrop sounds.
The rules of raindrops are that if a given number:
	- has 3 as a factor, add 'Pling' to the result.
	- has 5 as a factor, add 'Plang' to the result.
	- has 7 as a factor, add 'Plong' to the result.
	- does not have any of 3, 5, or 7 as a factor,
	  the result should be the digits of the number.
*/
func Convert(note int) string {
	var result string

	if note%3 == 0 {
		result += "Pling"
	}
	if note%5 == 0 {
		result += "Plang"
	}
	if note%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		return strconv.Itoa(note)
	}

	return result
}
