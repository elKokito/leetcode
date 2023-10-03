package main

import "fmt"

func main() {
  fmt.Println(addNumbers([]int{1,2,3,8,9, 9}, []int{3,4,5,5}))
}

func addNumbers(num1 []int,num2 []int) []int {
  numB := num1
  numS := num2
  if len(num1) < len(num2) {
    numB = num2
    numS = num1
  }
  result := make([]int, len(numB)+1, len(numB)+1)
  carry := 0
  for idx, _ := range(numS) {
    reminder := (numB[idx] + numS[idx]) % 10
    result[idx] = reminder + carry
    carry = (numB[idx] + numS[idx]) / 10
    fmt.Println(numB[idx], numS[idx])
    fmt.Println(reminder)
    fmt.Println(carry)
  }
  if len(num1) == len(num2) {
    result[len(num1)] = result[len(num1)] + carry
  } else {
    fmt.Println("here")
    for idx, _ := range numB[len(numS):] {
      result[len(numS) + idx] = carry + numB[len(numS) + idx]
      carry = 0
    }
  }
  return result
}
