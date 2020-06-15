package algorithms

import "testing"

func TestFizzbuzz(t *testing.T) {
	t.Run("0", testFizzbuzzFunc(0, "1"))
	t.Run("1", testFizzbuzzFunc(1, "1"))
	t.Run("3", testFizzbuzzFunc(3, "1 2 fizz"))
	t.Run("5", testFizzbuzzFunc(5, "1 2 fizz 4 buzz"))
	t.Run("15", testFizzbuzzFunc(15, "1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz"))
}

func testFizzbuzzFunc(num int, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := Fizzbuzz(num)
		if actual != expected {
			t.Errorf("Fizzbuzz(%d) => '%v'; actual: '%v'", num, actual, expected)
		}
	}
}
