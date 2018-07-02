package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

// Walk the tree t, sending all values from the tree to channel ch
func walkTree(t *tree.Tree, ch chan int){
  if t.Left != nil {
    walkTree(t.Left, ch)
  }
  ch <- t.Value
  if t.Right != nil {
    walkTree(t.Right,ch)
  }
}

func Walk(t *tree.Tree, ch chan int){
  walkTree(t, ch)
  close(ch)
}


func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for i := range ch1 {
    if i != <-ch2 {
      return false
    }
  }

  return true
}

func main() {
  test1 := Same(tree.New(1), tree.New(1)) // should return true
  test2 := Same(tree.New(1), tree.New(2)) // should return false
  fmt.Println("Test1 result:", test1, ", should be true")
  fmt.Println("Test2 result:", test2, ", should be false")
}
