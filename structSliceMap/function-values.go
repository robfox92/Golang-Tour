package main

import (
  "fmt"
  "math"
)

func compute (fn func(float64, float64) float64) float64 {
  // takes a function as input, and returns a float64
  //   the input function takes two float64 as input and
  //   returns a float64
  return fn(3,4)
}

func main() {
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(5,12))

  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))
}
