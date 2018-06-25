package main

import(
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

// method receiver appears in its own argument list between `func`
// and method name
// so in this case, Abs method has a receiver of type `Vertex` named `v`
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
}
