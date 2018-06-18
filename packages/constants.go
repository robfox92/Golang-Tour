package main

import "fmt"

// Constants are declared like variables
// but use the `const` keyword. They cannot be declared with `:=`
const Pi = 3.14

func main(){
  const World = "world!!"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
