package main

import (
  "fmt"
  "sync"
  "time"
)

type SafeCounter struct {
  v map[string]int
  mux sync.Mutex
}

func (c *SafeCounter) Inc(key string){
  // lock so that only one goroutine can access the value
  c.mux.Lock()
  c.v[key]++
  c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
  // defer the Unlock to ensure it happens
  c.mux.Lock()
  defer c.mux.Unlock()
  return c.v[key]
}

func main() {
  c := SafeCounter{v: make(map[string]int)}
  for i := 0; i < 1000; i++ {
    go c.Inc("somekey")
  }
  time.Sleep(time.Second)
  fmt.Println(c.Value("somekey"))
}
