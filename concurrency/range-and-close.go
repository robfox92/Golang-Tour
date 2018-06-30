package main

import (
  "fmt"
)

func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  // senders can close a channel to indicate that no more values will be
  // sent. Receivers can test by assigning `v, ok := <- ch` and testing
  // `ok`. If `ok == false`, there are no more values to receive.
  // usually you only need to do this when the receiver needs to know
  // that there will be no new values coming, e.g. when terminating a
  // range loop
  close(c)
}

func main() {
  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  for i := range c {
    fmt.Println(i)
  }
}
