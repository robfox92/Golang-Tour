package main

import(
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  // v is available to anything within the if/else block
  if v:= math.Pow(x,n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here as it is out of scope
  return lim
}

func main(){
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
