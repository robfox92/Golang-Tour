package main

import "fmt"

func add(x int, y int) int {
  // This function takes two ints as input and returns an int
  // Can also declare as:
  // func add(x, y int) int{}
  // When two or more consecutively named function parameters share
  // a type, you can omit the type from all but the last
  return x + y
}

func main(){
  fmt.Println(add(42,13))
}
