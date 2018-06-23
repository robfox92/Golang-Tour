package main

import "fmt"

func main() {
  fmt.Println("counting")

  // deferred function calls are pushed on to a stack
  // when the surrounding function returns, its deferred calls are
  // executed last-in-first-out
  for i:=0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("done")
}
