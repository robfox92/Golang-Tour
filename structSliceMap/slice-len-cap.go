package main

import "fmt"

func main() {
  s := []int{2,3,5,7,11,13}
  printSlice(s)

  // Slice the slice to give it 0 length
  s = s[:0]
  printSlice(s)

  // Extend the slice's length - you can only do this if the underlying
  // array has any room
  s = s[:4]
  printSlice(s)

  // Drop the first two values
  s = s[2:]
  printSlice(s)
}

func printSlice(s []int) {
  // length of the slice `len(s)` is the number of elements it contains
  // capacity of the slice `cap(s)` is the number of elements in the
  // underlying array, counting from the first element in the slice
  fmt.Printf("len=%d cap=%d val=%v\n",
len(s), cap(s), s)
}
