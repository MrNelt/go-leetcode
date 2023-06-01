package valid_sudoku

import (
	"fmt"
	"unicode"
)

func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		set1 := make(map[byte]bool)
		set2 := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' && (!unicode.IsDigit(rune(board[i][j])) || set1[board[i][j]]) {
				return false
			}
			if board[j][i] != '.' && (!unicode.IsDigit(rune(board[j][i])) || set2[board[j][i]]) {
				return false
			}
			set1[board[i][j]] = true
			set2[board[j][i]] = true
		}
	}
	for i := 0; i < 7; i += 3 {
		for j := 0; j < 7; j += 3 {
			fmt.Println(i, j)
			set := make(map[byte]bool)
			for i0 := i; i0 < i+3; i0++ {
				for j0 := j; j0 < j+3; j0++ {
					if board[i0][j0] != '.' && (!unicode.IsDigit(rune(board[i0][j0])) || set[board[i0][j0]]) {
						return false
					}
					set[board[i0][j0]] = true
				}
			}
		}
	}
	return true
}
