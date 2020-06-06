package algorithms

import "strconv"

// Fizzbuzz will, given an integer, print out all the numbers from 1 to that number
// If a number is divisible by 3, 'fizz' will be printed instead.
// If a number is divisible by 5, 'buzz' will be printed.
// If a number is divisible by 15, 'fizzbuzz' will be printed.
func Fizzbuzz(n int) string {
	output := "1"
	delimiter := " "

	if n <= 1 {
		return output
	}

	for i := 2; i <= n; i++ {
		num := strconv.Itoa(i)

		output += delimiter

		if i%15 == 0 {
			output += "fizzbuzz"
		} else if i%5 == 0 {
			output += "buzz"
		} else if i%3 == 0 {
			output += "fizz"
		} else {
			output = output + num
		}
	}

	return output
}
