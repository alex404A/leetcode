package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	matrix := build(m, n, &grid)
	return find(m, n, &matrix, &grid)
}

func find(m int, n int, ptr1 *[][]int, ptr2 *[][]int) int {
	matrix := *ptr1
	grid := *ptr2
	if matrix[m-1][n-1] == -1 {
		matrix[m-1][n-1] = min(find(m-1, n, &matrix, &grid), find(m, n-1, &matrix, &grid)) + grid[m-1][n-1]
	}
	return matrix[m-1][n-1]
}

func build(m int, n int, ptr *[][]int) [][]int {
	grid := *ptr
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}
	matrix[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		matrix[0][i] = matrix[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		matrix[i][0] = matrix[i-1][0] + grid[i][0]
	}
	return matrix
}

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func test() [][]int {
	m := 3
	matrix := make([][]int, m)
	matrix[0] = []int{1, 3, 1}
	matrix[1] = []int{1, 5, 1}
	matrix[2] = []int{4, 2, 1}
	// m := 1
	// matrix := make([][]int, m)
	// matrix[0] = []int{0}
	return matrix
}

func main() {
	grid := test()
	result := minPathSum(grid)
	fmt.Println(result)
}
