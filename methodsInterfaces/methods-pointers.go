package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Defining `Scale` as a function on `*Vertex` (i.e. a pointer
// to a Vertex). Scale can modify the value to which the receiver
// points. As methods often need to modify their receiver, pointer
// receivers are more common than value receivers, as value receivers
// mean that the method operates on a copy of the original value.
func  (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}



func main() {
  v:= Vertex{3,4}
  fmt.Println(v.Abs())

  v.Scale(10)
  fmt.Println(v.Abs())
}
