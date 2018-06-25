package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  // the `range` form of the `for` loop iterates over a slice or map.
  // when ranging over a slice, the index and a copy of the element
  // are returned
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}
