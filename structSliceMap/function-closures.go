package main

import "fmt"

func adder() func(int) int {
  // closures are functions that reference variables from outside their
  // body
  sum := 0
  return func(x int) int {
    sum +=x
    return sum
  }
}

func main() {
  pos, neg := adder(), adder()
  for i:= 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }
}
