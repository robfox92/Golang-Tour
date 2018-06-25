package main

import "fmt"

func main() {
  pow := make([]int, 10)

  // we only want the index here, so we don't assign the value
  for i := range pow {
    pow[i] = 1 << uint(i) // == 2**i
  }

  // we only want the value here, so we assign index to `_`
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}
