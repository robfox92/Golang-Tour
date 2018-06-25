package main

import(
  "fmt"
  "math"
)

// can only declare methods with a receiver whose type is defined
// in the same package as the method.
// you can't declare a method with a receiver whose type is defined
// in another package - including built-ins like int and float64
type MyFloat float64

// defining an Abs method for MyFloat type
func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func main() {
  f := MyFloat(-math.Sqrt(2))
  fmt.Println(f.Abs())
}
