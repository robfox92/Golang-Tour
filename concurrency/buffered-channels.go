package main

import "fmt"

func main() {
  // by providing a second argument to make(), we set the buffer length
  // this means the channel will only block from sends when the buffer is
  // full. Conversely, the channel will block from receives when the 
  // buffer is empty.
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
