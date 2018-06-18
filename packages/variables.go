package main

import "fmt"

// Package level variables - "globals" (??)
var c, python, java bool

func main(){
  // Function level variables
  var i int
  fmt.Println(i, c, python, java)
}
