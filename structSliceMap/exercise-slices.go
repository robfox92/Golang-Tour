package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  // returns a slice of length dy
  //   each element of this slice is a slice of length dx,
  //   each element of this slice is an 8-bit unsigned int

  // create the 2d-array - len(arr) == dy
  arr := make([][]uint8, dy)

  for i := range arr {
    // create the sub-arrays - lenn(arr[i]) == dx
    arr[i] = make([]uint8, dx)

    // assign values to the arrays
    for j:= range arr[i]{
      arr[i][j] = uint8((i+j)/2)
    }
  }
  return arr
}

func main() {
  pic.Show(Pic)
}
