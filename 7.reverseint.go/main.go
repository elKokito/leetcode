package main

import (
  "fmt"
  "strconv"
  "math"
)

func main() {
  fmt.Println(reverseInt(123))
  fmt.Println(reverseInt(987))
  fmt.Println(reverseInt(1234567))
}

func reverseInt(x int) int {
  negative := x < 0
  i := int(math.Abs(float64(x)))
  reverse := 0
  is := strconv.Itoa(i)
  numberDigit := len(is)
  pos := 0
  for pos < numberDigit{
    power := math.Pow(10, float64(numberDigit - 1 - pos))
    digit := i / int(power)
    reverse += digit * int(math.Pow(10, float64(pos)))
    i = i - digit * int(power)
    pos += 1
  }
  if reverse > int(math.Pow(2, 31)) {
    return 0
  }
  if negative {
    return reverse * -1
  } else {
    return reverse
  }
}
