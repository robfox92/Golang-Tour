package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

// m will map a string to a vertex
var m map[string]Vertex

func main() {

  // make returns a map of the given type, ready for assignment
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}
