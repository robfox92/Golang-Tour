package main

import "fmt"

func main() {
  // make a slice with len(5)
  a := make([]int, 5)
  printSlice("a", a)

  // make a slice with len(0) and cap(5)
  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  d := c[2:5]
  printSlice("d", d)
}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
