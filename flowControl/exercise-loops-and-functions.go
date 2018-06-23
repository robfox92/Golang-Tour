package main

import(
  "fmt"
)

func Sqrt(x float64) float64 {
  // Try to find some number z that satisfies z^2 - x = 0
  // initial estimate for z = 1
  z := float64(1);

  for i := 0; i < 10; i++ {
    z -= (z*z - x) / (2*z)
    fmt.Println(z)
  }
  return z
}

func main(){
  fmt.Println(Sqrt(2))
}
