package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  fmt.Println(v)
  p := &v
  p.X = 1e9        // don't need to explicitly dereference - e.g.
                   // (*p).X. Language permits implicit dereferencing.
  fmt.Println(v)
}
