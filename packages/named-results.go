package main

import "fmt"

func split(sum int) (x,y int){
  // We use named return values x and y
  // then, when we have a return statement ("naked return"),
  // we will return the named values. Should only be used
  // in short functions.
  x = sum * 4 / 9
  y = sum - x
  return
}

func main(){
  fmt.Println(split(17))
}
