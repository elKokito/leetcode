package main

import (
	"fmt"
	"math"
)

func main() {

  fmt.Println(getRainQuantity([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
  fmt.Println(getRainQuantity([]int{4,2,0,3,2,5}))
  fmt.Println(getRainQuantity([]int{4,2,3}))
  fmt.Println(getRainQuantity([]int{4,1,3,2}))
  fmt.Println(getRainQuantity([]int{5,2,1,3,3,2,2,4,1}))
}

func getRainQuantity(height []int) int {
  quantity := 0

  idx1 := 0
  // move to the first non-zero height
  for idx1 < len(height) && height[idx1] == 0 {
    idx1 += 1
  }

  idx2_max := 0
  ended := false
  for idx1 < len(height) {
    idx2 := idx1 + 1
    potentialWater := 0
    for idx2 < len(height) && height[idx2] < height[idx1] {
      idx2_max = int(math.Max(float64(idx2_max), float64(height[idx2])))
      potentialWater += height[idx1] - height[idx2]
      idx2 += 1
    }

    if idx2 < len(height) {
      quantity += potentialWater
      idx1 = idx2
    } else if !ended {
      height[idx1] = idx2_max
      ended = true
    } else {
      // just exiting, we could do "bresk" or idx1 = len(height)
      idx1 += 1
    }
  }

  return quantity
}
