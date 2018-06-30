package main

import (
  "fmt"
  "math"
)

type I interface {
  M()
}

type T struct {
  S string
}

func (t *T) M() {
  fmt.Println(t.S)
}

type F float64

func (f F) M() {
  fmt.Println(f)
}

func main() {
  var i I

  i = &T{"Hello"}
  describe(i)

  i = F(math.Pi)
  describe(i)
  i.M()

}

func describe(i I) {
  // Interface values are like a tuple of a value and a concrete type
  // e.g. (value, type)
  // Calling a method on an interface value executes the method of
  // the same name on its underlying type
  fmt.Printf("(%v, %T)\n", i, i)
}
