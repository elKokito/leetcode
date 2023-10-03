package main

import (
  "fmt"
  "strconv"
)

func main() {
  fmt.Println(isPalindrome(123))
  fmt.Println(isPalindrome(121))
  fmt.Println(isPalindrome(1))
}

func isPalindrome(number int) bool {
  snumber := strconv.Itoa(number)
  tail := len(snumber) -1
  head := 0
  for head < tail {
    if snumber[tail] != snumber[head] {
      return false
    }
    tail -= 1
    head += 1
  }
  return true
}
