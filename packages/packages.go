package main

// Package name is the same as the last element of the import path
// The "math/rand" package comprises of files that begin with the
// statement "package rand"

import(
  "fmt"
  "math/rand"
  "time"
)

func main(){
  currTime := time.Now().UnixNano() // get the current UNIX time
  rand.Seed(currTime)               // set as the random seed
  fmt.Println("My favourite number is",rand.Intn(10))
}
