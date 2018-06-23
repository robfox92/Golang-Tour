package main

import "fmt"

func main() {
  i := 42

  p := &i         // point to i
  fmt.Println(*p) // read i through the pointer

  // *p means to dereference p - it denotes the pointer's
  // underlying value
  *p = 21         // set i through the pointer
  fmt.Println(i)  // see the new value of i
  fmt.Println(p)  // see the pointer
//  q := p+1 //doesn't work - no pointer arithmetic
//  fmt.Println(*q)
}
