package main

import "fmt"

type Person struct {
  Name string
  Age int
}

// Stringer is an interface that describes types that can describe
// themselves as Strings. This means you need a String() method defined.
func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  a := Person{"Arthur Dent", 42}
  z := Person{"Zaphod Beeblebrox", 9001}
  fmt.Println(a, z)
}
