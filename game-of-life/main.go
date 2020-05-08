package main

import "fmt"

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			lives := getLiveNum(i, j, board)
			if board[i][j] == 1 && lives >= 2 && lives <= 3 {
				board[i][j] = 3
			}
			if board[i][j] == 0 && lives == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] >>= 1
		}
	}
}

func getLiveNum(row int, col int, board [][]int) int {
	lives := 0
	heigth := len(board)
	width := len(board[0])
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i < 0 || i >= heigth || j < 0 || j >= width {
				continue
			}
			if i == row && j == col {
				continue
			}
			if board[i][j] == 1 || board[i][j] == 3 {
				lives++
			}
		}
	}
	return lives
}

func main() {
	board := [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{0, 0, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}
