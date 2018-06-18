package main

import "fmt"

func main(){
  sum := 1
  // init and post omitted - this is what a while loop looks like
  // can also omit the semicolons e.g.:
  // `for sum < 1000`
  for ; sum < 1000 ;{
    sum += sum
  }
  fmt.Println(sum)
}
