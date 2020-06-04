package algorithms

// Sum will return the sum the integers in a slice.
func Sum(xi []int) int {

	sum := 0

	for _, num := range xi {
		sum += num
	}

	return sum
}
