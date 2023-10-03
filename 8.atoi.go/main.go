package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
  fmt.Println(atoi("    -00042"))
  fmt.Println(atoi("    -00042 with owr"))
  fmt.Println(atoi("-+12"))
}

func atoi(s string) int {
    output := 0

  index := 0
  if len(s) == 0 {
    return 0
  }
  // remove leading whitespace
  for s[index] == ' ' {
    index += 1
    if index == len(s) {
      return 0
    }
  }

  sign := 1
  // check if '+' or '-' character
  if s[index] == '-' {
    sign = -1
    index += 1
  } else if s[index] == '+' {
    index += 1
  }
  if len(s[index:]) == 0 {
    return 0
  }
  c := s[index]
  for c == '0' {
      index += 1
      if index < len(s) {
        c = s[index]
      } else {
        return 0
      }
  }
    if !strings.Contains("1234567890", string(s[index])) {
    return 0
  }

  alphabet := "abcdefghijklmnopqrstuvwxzyABCDEFGHIJKLMNOPQRSTUVWXYZ .,+-"
  end := strings.IndexAny(s[index:], alphabet)
  if end == -1 {
    end = len(s[index:])
  }
  mynumber := s[index:index+end]

  power := len(mynumber) - 1
  for _, number := range(mynumber) {
    switch number {
    case '1':
      output += int(1 * math.Pow(10, float64(power)))
    case '2':
      output += int(2 * math.Pow(10, float64(power)))
    case '3':
      output += int(3 * math.Pow(10, float64(power)))
    case '4':
      output += int(4 * math.Pow(10, float64(power)))
    case '5':
      output += int(5 * math.Pow(10, float64(power)))
    case '6':
      output += int(6 * math.Pow(10, float64(power)))
    case '7':
      output += int(7 * math.Pow(10, float64(power)))
    case '8':
      output += int(8 * math.Pow(10, float64(power)))
    case '9':
      output += int(9 * math.Pow(10, float64(power)))
    case '0':
    default:
      break
    }

    power -= 1
    index += 1
  }
    if math.Abs(float64(output)) >= math.Pow(2, 31) {
    output = int(math.Pow(2, 31) -1)
    if sign == -1 {
      output = output + 1
    } 
  }

  return output * sign
}
