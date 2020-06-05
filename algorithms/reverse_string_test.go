package algorithms

import "testing"

func TestReverseString(t *testing.T) {
	t.Run("foobar", testReverseStringFunc("foobar", "raboof"))
	t.Run("barfoo", testReverseStringFunc("barfoo", "oofrab"))
	t.Run("", testReverseStringFunc("", ""))
}

func testReverseStringFunc(input, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := ReverseString(input)
		if actual != expected {
			t.Errorf("ReverseString(%v) => %v, should be %v", input, actual, expected)
		}
	}
}
