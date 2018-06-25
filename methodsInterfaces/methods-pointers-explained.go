package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3,4}
  // Need to use a pointer here as we want to modify v.
  // alteranately, could use v = Scale(v, 10) and have Scale()
  // return a Vertex
  Scale(&v, 10)
  fmt.Println(Abs(v))
}
