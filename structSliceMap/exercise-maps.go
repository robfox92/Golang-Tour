package main

import(
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount (s string) map[string]int {
  // returns a map of the counts of each word in string s
  counts := make(map[string]int)
  for _, v := range strings.Fields(s) {
    counts[v]++
  }

  return counts
}

func main() {
  wc.Test(WordCount)
}
