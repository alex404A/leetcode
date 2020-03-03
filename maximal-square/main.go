package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	max := 0
	for i := 0; i+max < len(matrix); i++ {
		for j := 0; j+max < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				for length := max + 1; length <= len(matrix[0])-j; length++ {
					isSquare := findSquare(i, j, length, matrix)
					if isSquare {
						max = length
					} else {
						break
					}
				}
			}
		}
	}
	return max * max
}

func findSquare(row int, col int, length int, matrix [][]byte) bool {
	if row+length > len(matrix) || col+length > len(matrix[0]) {
		return false
	}
	for i := row; i < row+length; i++ {
		for j := col; j < col+length; j++ {
			if matrix[i][j] == '0' {
				return false
			}
		}
	}
	return true
}

func main() {
	matrix := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
	}
	fmt.Println(maximalSquare(matrix))
}
