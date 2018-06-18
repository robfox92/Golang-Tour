package main

import "fmt"

func main(){
  var i, j int = 1, 2
  // Inside a function we can use `:=`, the short assignment operator
  // It is used in place of a var declaration with an implicit type
  // It can only be used within a function, as every statement
  // outside of a function behins with a keyword so the `:=` construct
  // is not available
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}
