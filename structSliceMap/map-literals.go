package main

import "fmt"

type Vertex struct{
  Lat, Long float64
}

// map literals are like struct literals, but you need to provide a key
var m = map[string]Vertex{
  "Bell Labs": Vertex{
    40.68433, -74.39967,
  },
  "Google": Vertex{
    37.42202, -122.08408,
  },
}

func main() {
  fmt.Println(m)
}
