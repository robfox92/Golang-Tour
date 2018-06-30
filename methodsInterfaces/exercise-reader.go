package main

import "golang.org/x/tour/reader"

type MyReader struct {}

// TODO: Add a Read([]byte) (int,error) method to MyReader
func (m MyReader) Read([]byte) (int, error) {
  return 'A', nil
}

func main() {
  reader.Validate(MyReader{})
}
