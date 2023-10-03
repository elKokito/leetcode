package main

import "fmt"

func main() {
  fmt.Println(palindrome("babad"))
  fmt.Println(palindrome("babab"))
  fmt.Println(palindrome("aaabababbbbb"))
}

func palindrome(s string) string {
  maxPal := ""
  for idx, _ := range(s) {
    end := len(s)
    for end > idx {
      if isPalindrome(s[idx:end]) && (end - idx) > len(maxPal) {
        maxPal = s[idx:end]
      }
      end -= 1
    }
  }
  return maxPal
}

func isPalindrome(s string) bool {
  for i:= 0; i< len(s)/2 ; i++ {
    if s[i] != s[len(s) - i -1] {
      return false
    }
  }
  return true
}
