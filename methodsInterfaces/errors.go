package main

import (
  "fmt"
  "time"
)

// Go programs express error state with `error` values.
// The `error` type is a builtin interface similar to fmt.Stringer
// Functions often return an `error` value, and the calling function
// should handle errors by testing whether the error equals `nil`
// `error` == `nil` denotes success, `error` != `nil` denotes failure.
type MyError struct {
  When time.Time
  What string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error{
  return &MyError{
    time.Now(), "it didn't work",
  }
}

func main() {
  if err := run(); err != nil {
    fmt.Println(err)
  }
}
