package main

import "fmt"

func do(i interface{}) {
  // a type switch is a construct that permits several type assertions
  // in series. It is similar to regular switch statements.
  switch v := i.(type) {
  // In this case, we test to see if an interface value `i` is an `int` or
  // a `string`.
  // In the `case int` case, the variable `v` has type `int`. In the
  // `default` case, the variable `v` has the same type and value as `i`
  case int:
    fmt.Printf("twice %v is %v\n",v,v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}

func main() {
  do(21)
  do("hello")
  do(true)
}
