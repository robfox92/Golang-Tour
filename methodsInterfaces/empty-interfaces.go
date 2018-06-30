package main

import "fmt"

func main() {
  // declare i as an empty interface - this specifies 0 methods.
  // it may hold values of any type, as every type implements at least 0
  // methods.

  // Empty interfaces are used by code that handles values of unknown type
  // for example, fmt.Print takes any number of arguments of type
  // interface{}
  var i interface{}
  describe(i)

  i = 42
  describe(i)

  i = "hello"
  describe(i)
}

func describe(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}
