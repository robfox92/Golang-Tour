package main

import "fmt"

type Vertex struct {
  X, Y int
}

var(
  v1 = Vertex{1,2}      // has type Vertex
  v2 = Vertex{X: 1}     // Y:0 is implicit. Also assigning variables using
                        // Name: syntax. In this case, order does not
                        // matter due to naming.
  v3 = Vertex{}         // X:0 and Y:0 is implicit
  p = &Vertex{1,2}      // has type *Vertex
                        // The prefix `&` returns a pointer to the struct
                        // value
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
