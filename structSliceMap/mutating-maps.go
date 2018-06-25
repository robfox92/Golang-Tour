package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The answer:", m["Answer"])

  m["Answer"] = 7
  fmt.Println("The answer:", m["Answer"])

  delete(m, "Answer")
  // if key not in map, returned element is the zero value for the
  // map's element type
  fmt.Println("The answer:", m["Answer"])

  // test a value is present - if key in map, ok == true
  v, ok := m["Answer"]
  fmt.Println("The answer:", v, "Present?", ok)

  m["Answer"] = 42
  v, ok = m["Answer"]
  fmt.Println("The answer:", v, "Present?", ok)

}
