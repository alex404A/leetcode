package main

import "fmt"

func setZeroes(matrix [][]int) {
	first := iter(&matrix)
	replace(&matrix, first)
}

func iter(ptr *[][]int) int {
	matrix := *ptr
	first := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
				if i == 0 && j != 0 {
					if first == 3 || first == 4 {
						first = 4
					} else {
						first = 2
					}
				} else if i != 0 && j == 0 {
					if first == 2 || first == 4 {
						first = 4
					} else {
						first = 3
					}
				} else if i == 0 && j == 0 {
					first = 4
				}
			}
		}
	}
	return first
}

func replace(ptr *[][]int, first int) {
	matrix := *ptr
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			replaceRow(i, ptr)
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			replaceCol(j, ptr)
		}
	}
	if first == 2 {
		replaceRow(0, ptr)
	} else if first == 3 {
		replaceCol(0, ptr)
	} else if first == 4 {
		replaceRow(0, ptr)
		replaceCol(0, ptr)
	}
}

func replaceRow(row int, ptr *[][]int) {
	matrix := *ptr
	for j := 0; j < len(matrix[0]); j++ {
		matrix[row][j] = 0
	}
}

func replaceCol(col int, ptr *[][]int) {
	matrix := *ptr
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}

type point struct {
	x int
	y int
}

func build(m int, n int, points []point) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = 1
		}
	}
	for _, point := range points {
		matrix[point.x][point.y] = 0
	}
	return matrix
}

func main() {
	points := make([]point, 6)
	points[0] = point{0, 0}
	points[1] = point{0, 1}
	points[2] = point{0, 2}
	points[3] = point{2, 0}
	points[4] = point{4, 0}
	points[5] = point{4, 1}
	matrix := build(5, 4, points)
	setZeroes(matrix)
	fmt.Println("finish")
}
