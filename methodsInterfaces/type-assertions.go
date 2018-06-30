package main

import "fmt"

func main() {
  var i interface{} = "hello"


  // A type assertion provides access to an interface value's underlying
  // concrete value

  s := i.(string) // asserts that `i` holds the concrete type `string`,
                  // and assigns the underlying string value to `s`
  fmt.Println(s)

  s, ok := i.(string) // tests whether `i` holds a string, stores the
                      // result of that test in `ok`. If `i` holds a
                      // string, then i will be the underlying value and
                      // `ok` will be `true`
  fmt.Println(s, ok)

  f, ok := i.(float64)// if `i` does not hold a `float64`, then `ok`
  fmt.Println(f, ok)  // will be `false` and `f` will have the zero value
                      // of `float64`

//  f = i.(float64)   // triggers a panic as `i` is not of type `float64`
//  fmt.Println(f)
}
