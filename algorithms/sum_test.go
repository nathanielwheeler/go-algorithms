package algorithms

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("[1,2,3,4,5]", testSumFunc([]int{1, 2, 3}, 6))
	t.Run("[1,2,3,4,5]", testSumFunc([]int{1, 2, 3, -6}, 0))
	t.Run("[1,2,3,4,5]", testSumFunc([]int{-1, -2, -3}, -6))
	t.Run("[]", testSumFunc([]int{}, 0))
	t.Run("nil", testSumFunc(nil, 0))
}

func testSumFunc(xi []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := Sum(xi)
		if actual != expected {
			t.Errorf("Sum(%v) => %d, should be %d", xi, actual, expected)
		}
	}
}
