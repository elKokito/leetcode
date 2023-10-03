package main

import "fmt"

func main() {
  fmt.Println("ab" == "ab")
  fmt.Println("ab " == "abc")
  fmt.Println(isMatch("ab", "ab"))
  fmt.Println(isMatch("ab", "a.*"))
  fmt.Println(isMatch("alaksdjf", "al.*f"))
}

func isMatch(s string, pattern string) bool {

  previous := pattern[0]
  if previous == '.' {
  }
  if previous == '*' {
    return false
  }
  if previous != s[0] {
    return false
  }
  s_index := 1
  index := 1
  for index < len(pattern) {
    switch pattern[index] {
    case '.':
      previous = '.'
      s_index += 1
      index += 1
    case '*':
      if previous == '.' {
        s_index = len(s) - 1
        index += 1
      } else {
        for s[s_index] == previous {
          s_index += 1
        }
      }
    default:
      previous = s[s_index]
      s_index += 1
      index += 1
    }
  }
  if index - 1 == len(pattern) - 1 && s_index != len(s) - 1 || s_index == len(s) - 1 && index -1 != len(pattern) - 1 { return false }
  return true
}
