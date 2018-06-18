package main

import "fmt"

const(
  // Create a large number by left shifting 100 places
  Big = 1 << 100
  // Shift it right 99 places so we have the number 2
  Small = Big >> 99
)

func needInt(x int) int {
  return x*10+1
}
func needFloat(x float64) float64{
  return x * 0.1
}

func main(){
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}
