package main

import "fmt"

func main(){
  sum := 0
  // `for` is the only looping construct
  // syntax:
  //    `for <init>;<condition>;<post>`
  // init - (optional) statement run prior to the first iteration
  // condition - (optional) evaluated before every iteration
  // post - (optional) statement executed at the end of every iteration

  for i:=0; i<10; i++{
    sum += i
  }
  fmt.Println(sum)
}
