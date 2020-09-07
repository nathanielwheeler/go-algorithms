package algorithms

import (
  "fmt"
)

// Hamming will take in a byte of up to 16 bits and check it for  up to 1 error, returning the location of that error.
// TODO Should be able to be generalized to byte lengths of multiples of 16
func Hamming(input uint16) uint {
  // Turn input into a binary string
  bitstr := fmt.Sprintf("%b", input)

  // Make a slice of bit locations
  var bitLoc []uint 
  for i, v := range bitstr {
    if string(v) == "1" {
      bitLoc = append(bitLoc, uint(i))
    }
  }
  fmt.Println(bitLoc)

  // XOR the bit locations
  loc1 := bitLoc[0]
  for i := 1; i <= len(bitLoc[1:]); i++ {
    loc2 := bitLoc[i]
    fmt.Printf("%d ^ %d = %d\n", loc1, loc2, loc1 ^ loc2)
    loc1 = loc1 ^ loc2
  }
  errorPosition := loc1
  return errorPosition
}