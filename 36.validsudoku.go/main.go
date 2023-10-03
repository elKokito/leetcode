package main

import "fmt"

func main() {
  board1 := [][]byte{
    {'1','2','3', '4','5','6', '7','8','9'},
    {'.', '4', '.','.','.','.','.','.','.'},
    {'.', '4', '.','.','.','.','.','.','.'},
    {'.', '4', '.','.','.','.','.','.','.'},
    {'.', '4', '.','.','.','.','.','.','.'},
    {'.', '4', '.','.','.','.','.','.','.'},
    {'.', '4', '.','.','.','.','.','.','.'},
    {'.', '4', '.','.','.','.','.','.','.'},
    {'.', '4', '.','.','.','.','.','.','.'},
  }
  board2 := [][]byte{
    {'1','2','3', '4','5','6','7','8','9'},
    {'.', '.', '.','.','.','.','.','.','.'},
    {'.', '.', '.','.','.','.','.','.','.'},
    {'.', '.', '.','.','.','.','.','.','.'},
    {'.', '.', '.','.','.','.','.','.','.'},
    {'.', '.', '.','.','.','.','.','.','.'},
    {'.', '.', '.','.','.','.','.','.','.'},
    {'.', '.', '.','.','.','.','.','.','.'},
    {'.', '.', '.','.','.','.','.','.','.'},
  }
  board3 := [][]byte{
{'5','3','.','.','7','.','.','.','.'},
{'6','.','.','1','9','5','.','.','.'},
{'.','9','8','.','.','.','.','6','.'},
{'8','.','.','.','6','.','.','.','3'},
{'4','.','.','8','.','3','.','.','1'},
{'7','.','.','.','2','.','.','.','6'},
{'.','6','.','.','.','.','2','8','.'},
{'.','.','.','4','1','9','.','.','5'},
{'.','.','.','.','8','.','.','7','9'},
  }

  board4 := [][]byte {
    {'.','.','4','.','.','.','6','3','.'},
    {'.','.','.','.','.','.','.','.','.'},
    {'5','.','.','.','.','.','.','9','.'},
    {'.','.','.','5','6','.','.','.','.'},
    {'4','.','3','.','.','.','.','.','1'},
    {'.','.','.','7','.','.','.','.','.'},
    {'.','.','.','5','.','.','.','.','.'},
    {'.','.','.','.','.','.','.','.','.'},
    {'.','.','.','.','.','.','.','.','.'},
  }
  fmt.Println(isValid(board1))
  fmt.Println(isValid(board2))
  fmt.Println(isValid(board3))
  fmt.Println(isValid(board4))
}

func isValid(board [][]byte) bool {

  for r:=0; r < len(board); r++ {
    if !isRowValid(board[r]) {
      return false
    }

    column := []byte{}
    for i := 0; i < 9; i++ {
      column = append(column, board[i][r])
    }
    if !isColumnValid(column) {
      return false
    }
  }
  for c := 0; c < 3; c++ {
    for r := 0; r < 3; r++ {
      s := [][]byte{board[3*r][3*c:3*c+3],board[3*r+1][3*c:3*c+3], board[3*r+2][3*c:3*c+3]}
      if !isSquareValid(s) {
        fmt.Println(s)
        return false
      }
    }
  }
  return true
}

func isRowValid(row []byte) bool {
  digit := map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8':false, '9':false}
  for _, value := range(row) {
    if value != '.' {
      if !digit[value] {
        digit[value] = true
      } else {
        return false
      }
    }
  }
  return true
}

func isColumnValid(column []byte) bool {
  return isRowValid(column)
}

func isSquareValid(square [][]byte) bool {
  digit := map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8': false, '9': false}
  for i := 0; i < 3; i += 1 {
    for j := 0; j < 3; j += 1 {
      if square[i][j] != '.' {
        if !digit[square[i][j]] {
          digit[square[i][j]] = true
        } else {
          return false
        }
      }
    }
  }
  return true
}
