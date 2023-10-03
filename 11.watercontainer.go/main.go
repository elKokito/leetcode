package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(maxWaterContainer([]int{1,2,3,4}))
  fmt.Println(maxWaterContainer([]int{1,8,6,2,5,4,8,3,7}))
  fmt.Println(maxWaterContainer([]int{1,1}))
}

func maxWaterContainer(container []int) int {
  maxVolume := 0

  if len(container) == 2 {
    return int(math.Min(float64(container[0]), float64(container[1])))
  }

  i := 0
  t := len(container) - 1
  for i < t {
    h1 := container[i]
    h2 := container [t]
    volume := (t-i) * int(math.Min(float64(h1), float64(h2)))
    maxVolume = int(math.Max(float64(maxVolume), float64(volume)))
    if h1 < h2{
      i += 1
    } else {
      t -= 1
    }
  }
  return maxVolume
}
