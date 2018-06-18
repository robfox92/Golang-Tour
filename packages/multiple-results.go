package main

import "fmt"

func swap(x, y string) (string, string){
  return y, x
}

func main(){
  a := "Hello"
  b := "World"
  c, d := swap(a,b)
  fmt.Println("Arguments:",a,",",b)
  fmt.Println("Results:",c,",",d)
}
