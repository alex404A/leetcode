package main

import "fmt"

func uniquePaths(m int, n int) int {
	matrix := build(m, n)
	return find(m, n, &matrix)
}

func find(m int, n int, ptr *[][]int) int {
	matrix := *ptr
	if matrix[m-1][n-1] == 0 {
		matrix[m-1][n-1] = find(m-1, n, &matrix) + find(m, n-1, &matrix)
	}
	return matrix[m-1][n-1]
}

func build(m int, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		matrix[0][i] = 1
	}
	for i := 0; i < m; i++ {
		matrix[i][0] = 1
	}
	return matrix
}

func main() {
	result := uniquePaths(3, 2)
	fmt.Println(result)
}
