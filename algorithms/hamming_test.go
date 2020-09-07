package algorithms

import (
  "testing"
)

const(
  t1 uint16 = 35446 // 1000101001110110
)

func TestHamming(t *testing.T) {
  t.Run("35446", testHammingFunc(t1, 9))
}

func testHammingFunc(input uint16, expected uint) func(*testing.T) {
  return func(t *testing.T) {
    actual := Hamming(input)
    if actual != expected {
      t.Errorf("Hamming(%v) => %b, instead: %b", input, actual, expected)
    }
  }
}

