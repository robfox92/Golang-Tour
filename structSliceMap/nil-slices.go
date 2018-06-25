package main

import "fmt"

func main() {
  var s []int

// s is a nil slice - length and capacity 0, and no underlying array
fmt.Println(s, len(s), cap(s))

  if s == nil {
    fmt.Println("nil!")
  }
}
