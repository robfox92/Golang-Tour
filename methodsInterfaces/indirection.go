package main

import "fmt"

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3,4}
  // Even though we are calling Scale() on a Vertex (not a *Vertex),
  // Go interprets v.Scale(2) as (&v).Scale(2) since Scale() has a
  // pointer receiver
  v.Scale(2)
  ScaleFunc(&v, 10)

  p := &Vertex{4,3}
  p.Scale(3)
  ScaleFunc(p,8)

  fmt.Println(v,p)
}
