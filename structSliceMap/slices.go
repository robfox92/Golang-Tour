package main

import "fmt"

func main() {
  primes := [6]int{2,3,5,7,11,13}

  var s []int = primes[1:4]      // declare a slice of ints
  fmt.Println(s)                 // a slice is a dynamically sized
                                 // flexible view of an array
                                 // slices are more common than arrays
}
