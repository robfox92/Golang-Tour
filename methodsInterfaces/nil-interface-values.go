package main

import "fmt"

type I interface {
  M()
}

func main() {
  var i I
  describe(i)
  // calling a method on a nil interface causes a runtime error
  // as there is no type inside the interface tuple to indicate which
  // concrete method to call
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}
