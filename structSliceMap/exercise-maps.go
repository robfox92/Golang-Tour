package main

import(
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount (s string) map[string]int {
  // returns a map of the counts of each word in string s
  counts := make(map[string]int)
  fields := strings.Fields(s)
  for _, v := range fields {
    counts[v]++
  }

  return counts
}

func main() {
  wc.Test(WordCount)
}
