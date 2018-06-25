package main

import "fmt"

// fibonacci is a function that returns a function that returns an int
// the closure returns successive fibonacci numbers
func fibonacci() func() int {
  a := 0
  b := 1
  c := 0
  return func() int {
    c, a, b = a, b, a+b
    return c
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
