package main

import "fmt"

type Location struct {
	row int
	col int
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	outer := iter(board)
	visited := make(map[Location]bool)
	connected := make(map[Location]bool)
	border := Location{len(board) - 1, len(board[0]) - 1}
	for location := range outer {
		if _, ok := visited[location]; !ok {
			dfs(location, board, visited, connected, border)
		}
	}
	for i := 1; i < border.row; i++ {
		for j := 1; j < border.col; j++ {
			if board[i][j] == 'O' {
				if _, ok := connected[Location{i, j}]; !ok {
					board[i][j] = 'X'
				}
			}
		}
	}
}

func iter(board [][]byte) (outer map[Location]bool) {
	outer = make(map[Location]bool)
	height := len(board)
	width := len(board[0])
	for i, row := range board {
		for j := range row {
			if (i == 0 || i == height-1 || j == 0 || j == width-1) && board[i][j] == 'O' {
				outer[Location{i, j}] = true
			}
		}
	}
	return
}

func dfs(start Location, board [][]byte, visited map[Location]bool, connected map[Location]bool, border Location) {
	visited[start] = true
	connected[start] = true
	neighbours := getAdjacent(start, board, connected, border)
	for _, neighbour := range neighbours {
		if _, ok := connected[neighbour]; ok {
			continue
		}
		if neighbour.row == 0 || neighbour.row == border.row || neighbour.col == 0 || neighbour.col == border.col {
			visited[neighbour] = true
		}
		connected[neighbour] = true
		dfs(neighbour, board, visited, connected, border)
	}
}

func getAdjacent(start Location, board [][]byte, connected map[Location]bool, boader Location) []Location {
	neighbours := make([]Location, 0)
	if start.row > 0 {
		up := Location{start.row - 1, start.col}
		if _, ok := connected[up]; !ok && board[up.row][up.col] == 'O' {
			neighbours = append(neighbours, up)
		}
	}
	if start.col > 0 {
		left := Location{start.row, start.col - 1}
		if _, ok := connected[left]; !ok && board[left.row][left.col] == 'O' {
			neighbours = append(neighbours, left)
		}
	}
	if start.row < boader.row {
		down := Location{start.row + 1, start.col}
		if _, ok := connected[down]; !ok && board[down.row][down.col] == 'O' {
			neighbours = append(neighbours, down)
		}
	}
	if start.col < boader.col {
		right := Location{start.row, start.col + 1}
		if _, ok := connected[right]; !ok && board[right.row][right.col] == 'O' {
			neighbours = append(neighbours, right)
		}
	}
	return neighbours
}

func main() {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)
	// board := [][]byte{
	// 	[]byte{'O', 'O'},
	// 	[]byte{'O', 'O'},
	// }
	// solve(board)
	// fmt.Println(board)
	// board := [][]byte{
	// 	[]byte{'O', 'X', 'X', 'X'},
	// 	[]byte{'O', 'O', 'X', 'X'},
	// 	[]byte{'X', 'X', 'O', 'O'},
	// 	[]byte{'X', 'O', 'X', 'O'},
	// 	[]byte{'X', 'X', 'X', 'X'},
	// }
	// solve(board)
	// fmt.Println(board)
}
