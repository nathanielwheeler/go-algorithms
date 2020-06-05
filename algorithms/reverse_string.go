package algorithms

// ReverseString will take in a string and reverse the characters in that string.
func ReverseString(str string) string {
	output := ""

	for i := len(str) - 1; i >= 0; i-- {
		char := string(str[i])
		output += string(char)
	}

	return output
}
