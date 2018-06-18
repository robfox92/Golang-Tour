package main

import(
  "fmt"
  "math"
)

func main(){
  // A name is exported from a package if it begins with a capital letter
  // so math.pi is not available to us, but math.Pi is
  // when importing a package you can only refer to the exported names
  // unexported names are obviously not available from outside the package
  fmt.Println(math.Pi)
}
