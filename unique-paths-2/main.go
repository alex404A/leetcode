package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	} else if len(obstacleGrid[0]) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	build(m, n, &obstacleGrid)
	return find(m, n, &obstacleGrid)
}

func find(m int, n int, ptr *[][]int) int {
	matrix := *ptr
	if matrix[m-1][n-1] == -1 {
		return 0
	}
	if matrix[m-1][n-1] == 0 {
		matrix[m-1][n-1] = find(m-1, n, &matrix) + find(m, n-1, &matrix)
	}
	return matrix[m-1][n-1]
}

func build(m int, n int, ptr *[][]int) {
	obstacleGrid := *ptr
	initObstacle(m, n, ptr)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}
}

func initObstacle(m int, n int, ptr *[][]int) {
	obstacleGrid := *ptr
	if obstacleGrid[0][0] == 1 {
		for i := 0; i < m; i++ {
			obstacleGrid[i][0] = -1
		}
		for i := 0; i < n; i++ {
			obstacleGrid[0][i] = -1
		}
		return
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			for j := i; j < m; j++ {
				obstacleGrid[j][0] = -1
			}
			break
		} else {
			obstacleGrid[i][0] = 1
		}
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			for j := i; j < n; j++ {
				obstacleGrid[0][j] = -1
			}
			break
		} else {
			obstacleGrid[0][i] = 1
		}
	}
}

func test() [][]int {
	m := 1
	n := 1
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	matrix[0][0] = 1
	return matrix
}

func main() {
	obstacleGrid := test()
	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(result)
}
