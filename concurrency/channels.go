package main

import "fmt"

// Channels are typed conduits through which you can send/receive values
// with the channel operator `<-`. Data flows in the direction of the
// arrow. `ch <- v` means to send v to channel ch. `v := <-ch` means to
// receive from channel ch and assign the value to v.
//
// Sends and receives block until the other side is ready, providing
// synchronization without explicit locks

func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum //send sum to c
}

func main() {
  s := []int{7, 2, 8. -9, 4, 0}

  c := make(chan int)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)

  x, y := <-c, <-c //receive from c

  fmt.Println(x, y, x+y)
}
