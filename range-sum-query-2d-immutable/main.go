package main

import "fmt"

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{sum}
	}
	for i := 0; i < len(matrix); i++ {
		sum = append(sum, make([]int, len(matrix[0])))
	}
	sum[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		sum[i][0] = sum[i-1][0] + matrix[i][0]
	}
	for j := 1; j < len(matrix[0]); j++ {
		sum[0][j] = sum[0][j-1] + matrix[0][j]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sum[row2][col2]
	} else if row1 == 0 {
		return this.sum[row2][col2] - this.sum[row2][col1-1]
	} else if col1 == 0 {
		return this.sum[row2][col2] - this.sum[row1-1][col2]
	} else {
		return this.sum[row2][col2] - this.sum[row1-1][col2] - this.sum[row2][col1-1] + this.sum[row1-1][col1-1]
	}
}

func main() {
	matrix := [][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	fmt.Println(obj.SumRegion(0, 0, 3, 3))
	fmt.Println(obj.SumRegion(0, 1, 3, 3))
	fmt.Println(obj.SumRegion(1, 0, 3, 3))
	fmt.Println(obj.SumRegion(1, 1, 3, 3))
}
