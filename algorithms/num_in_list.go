package algorithms

import (
	// "fmt"
)

// NumInList will return true if num is in the list slice.
func NumInList(list []int, num int) bool {
	// fmt.Println("Hello from NumInList!")

	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
