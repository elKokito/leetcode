package main

import "fmt"

func main() {
  fmt.Println(median([]int{1,2,3} ,[]int{4,5,6,7,7,7,7}))
}

func median(arr1 []int, arr2 []int) float64 {
  merged := []int{}

  idx1 := 0
  idx2 := 0

  for idx1 < len(arr1) && idx2 < len(arr2) {
    if arr1[idx1] < arr2[idx2] {
      merged = append(merged, arr1[idx1])
      idx1 += 1
    } else {
      merged = append(merged, arr2[idx2])
      idx2 += 1
    }
  }
  if idx1 == len(arr1) {
    for i := idx2; i < len(arr2); i++ {
      merged = append(merged, arr2[i])
    }
  } else {
    for i := idx1; i < len(arr1); i++ {
      merged = append(merged, arr1[i])
    }
  }

  fmt.Println(merged)
  middleIndex := len(merged) / 2
  parity := len(merged) % 2 == 0
  if !parity {
    return float64(merged[middleIndex])
  } else {
    fmt.Println("here")
    fmt.Println(merged[middleIndex-1])
    fmt.Println(merged[middleIndex])
    val := float64(merged[middleIndex-1] + merged[middleIndex])
    return val / 2.0
  }
}
