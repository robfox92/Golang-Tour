package main

import "fmt"

type I interface {
  M()
}

type T struct{
  S string
}

// This method means type T implements interface I
// but we don't need to explcitly declare that it does so.
// Simply giving T the required methods (in this case, M()) means that
// T implements I

func (t T) M() {
  fmt.Println(t.S)
}

func main() {
  var i I = T{"hello"}
  i.M()
}
