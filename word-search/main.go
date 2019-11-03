package main

import "fmt"

type container struct {
	board    [][]byte
	searched [][]bool
}

type coordinate struct {
	x int
	y int
}

func (container *container) find(word string, ptr *coordinate) bool {
	c := *ptr
	if word[0] == container.board[c.x][c.y] && !container.searched[c.x][c.y] {
		if len(word) == 1 {
			return true
		}
		container.searched[c.x][c.y] = true
		if c.y-1 >= 0 {
			next := coordinate{c.x, c.y - 1}
			result := container.find(word[1:], &next)
			if result {
				return true
			}
		}
		if c.x-1 >= 0 {
			next := coordinate{c.x - 1, c.y}
			result := container.find(word[1:], &next)
			if result {
				return true
			}
		}
		if c.y+1 < len(container.board[0]) {
			next := coordinate{c.x, c.y + 1}
			result := container.find(word[1:], &next)
			if result {
				return true
			}
		}
		if c.x+1 < len(container.board) {
			next := coordinate{c.x + 1, c.y}
			result := container.find(word[1:], &next)
			if result {
				return true
			}
		}
		container.searched[c.x][c.y] = false
	}
	return false
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	searched := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		searched[i] = make([]bool, len(board[0]))
	}
	container := container{board, searched}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			result := container.find(word, &coordinate{i, j})
			if result {
				return true
			}
		}
	}
	return false
}

func test(board [][]byte, word string) {
	result := exist(board, word)
	fmt.Println(result)
}

func main() {
	board := make([][]byte, 3)
	board[0] = []byte{'A', 'B', 'C', 'E'}
	board[1] = []byte{'S', 'F', 'C', 'S'}
	board[2] = []byte{'A', 'D', 'E', 'E'}
	// test(board, "ABCCED")
	test(board, "SEE")
	test(board, "ABCB")
}
