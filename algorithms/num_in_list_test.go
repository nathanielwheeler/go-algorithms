package algorithms

import (
	"testing"
)

// TestNumInList will test if NumInList works
func TestNumInList(t *testing.T) {
	t.Run("[1,2,3,4,5]", testNumInListFunc([]int{1,2,3,4,5}, 5, true))
	t.Run("[1,2,3,4,5]", testNumInListFunc([]int{1,2,3,4,5}, 0, false))
}

func testNumInListFunc(xi []int, i int, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		actual := NumInList(xi, i)
		if actual != expected {
			t.Errorf("NumInList(%v, %d) => %t, should be %t", xi, i, actual, expected)
		}
	}
}