package main

import "fmt"

func main() {
  // defer means the function will not be executed until after
  // the surrounding function returns
  // in this case, the below line will
  //  - have its arguments evaluated immediately
  //  - only be executed after main() returns
  defer fmt.Println("world")

  fmt.Println("Hello")
}
