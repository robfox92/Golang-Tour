package main

import "fmt"

type I interface {
  M()
}

type T struct {
  S string
}

func (t *T) M() {
  if t == nil {
    fmt.Println("<nil>")
    return
  }
  fmt.Println(t.S)
}

func main() {
  var i I
  var t *T
  // if the concrete value inside the interface is nil, the method will
  // be called with a nil receiver. It is common to write methods that
  // gracefully handle being called with a nil receiver.
  // Note that an interface value that holds a nil concrete value is
  // itself non-nil.
  i = t
  describe(i)
  i.M()

  i = &T{"Hello"}
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}
