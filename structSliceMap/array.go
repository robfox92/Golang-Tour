package main

import "fmt"

func main() {
  var a [2]string  // declares an array of two strings
                   // array length is part of its type, cannot be resized
  a[0] = "Hello"
  a[1] = "World"

  fmt.Println(a[0],a[1])
  fmt.Println(a)

  primes := [6]int{2,3,5,7,11,13}
  fmt.Println(primes)
}
