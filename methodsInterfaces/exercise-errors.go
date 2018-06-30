package main

import (
  "fmt"
)

func Sqrt(x float64) (float64, error) {
  if x < 0{
    return nil, 1
  }
  z := float64(1)
  for i := 0; i < 10; i++ {
      z -= (z*z - x)/(2*z)
  }
  return z, nil
}

// TODO: implement ErrNegativeSqrt and ErrNegativeSqrt.Error()

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
