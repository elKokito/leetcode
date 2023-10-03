package main

import (
  "fmt"
  "strings"
)

func main() {
  sub := substring("abcabcaefg")
  fmt.Println(sub)
}

func substring(s string) int {
  max := 1
  currentMax := 0
  currentSub := ""

  for _, c := range(s) {
    if !strings.Contains(currentSub, string(c)) {
      currentSub += string(c)
      currentMax += 1
      if currentMax > max {
        max = currentMax
      }
    }
  }
  return max
}

