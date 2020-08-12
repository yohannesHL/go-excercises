
// Package twofer provides function ShareWith 
package twofer

import "fmt"


// ShareWith Returns a personalised message: One for x, one for me. 
// Where x is "you" If name empty otherwise will use given name
func ShareWith(name string) string {
	var username string = "you"
	if name != "" {
		username = name
	}
	return fmt.Sprintf("One for %s, one for me.", username)
}
