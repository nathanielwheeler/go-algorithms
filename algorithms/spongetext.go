package algorithms

import "strings"

// Spongetext will, given a string, turn it into sPoNgEtExT and return the result.
/* To do this, I need to ignore spaces and special characters.*/
func Spongetext(input string) string {
	output := ""

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if i % 2 == 0 {
			output += strings.ToLower(char)
		} else {
			output += strings.ToUpper(char)
		}
	}

	return output
}