package main

import (
  "fmt"
  "math"
)

type Abser interface {
  // Abser types implement the Abs() function, returning a float64
  Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

type Vertex struct{
  X, Y float64
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  // `a` can only be defined as things that implement Abser
  // i.e. anything that means that `a.Abs()` returns a float64
  var a Abser
  f := MyFloat(-math.Sqrt(2))
  v := Vertex{3,4}

  a = f  // a MyFloat implements Abser - MyFloat.Abs() is defined
  a = &v // a *Vertex also implements Abser - (*Vertex).Abs() is defined
//  a = v  // a Vertex does /not/ implement Abser, so this throws an error
  fmt.Println(a.Abs())
}
