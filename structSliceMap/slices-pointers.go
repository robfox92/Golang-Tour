package main

import "fmt"

func main() {
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println(names)

  // a and b don't really store any data - just descibe a section
  // of an underlying array
  a := names[0:2]
  b := names[1:3]
  fmt.Println(a,b)

  // other slices that share the same underlying array will see the
  // changes made to the slices
  b[0] = "XXX"
  fmt.Println(a,b)
  fmt.Println(names)
}
