package main

import "fmt"

func main() {
  numList := []int { 5, 10 , 15, 4}
  sol := twoSum( numList, 14)
  if sol != nil {
    fmt.Println(sol)
  } else {
    fmt.Println("no solution")
  }

  sol2 := twoSum2(numList, 14)
  if sol != nil {
    fmt.Println(sol2)
  }
}

func twoSum(nums []int, target int) []int {
  m := make(map[int]int)
  for idx, num := range nums {
    if val, found := m[target-num]; found {
      return []int{val, idx}
    }
    m[num] = idx
  }
  return nil
}

func twoSum2(nums []int, target int) []int {
  for idx, num := range nums {
    for idxx, numm := range nums[idx:] {
      if num + numm == target {
        return []int { idx, idx + idxx}
      }
    }
  }
  return []int {}
}
