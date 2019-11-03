package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	width := len(matrix[0]) - 1
	low := 0
	high := len(matrix) - 1
	row := -1
	for low <= high {
		mid := (low + high) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			if matrix[mid][width] >= target {
				row = mid
				break
			} else {
				low = mid + 1
			}
		} else if matrix[mid][0] > target {
			if mid > 0 && matrix[mid-1][0] <= target {
				row = mid - 1
				break
			} else {
				high = mid - 1
			}
		}
	}
	if row == -1 {
		return false
	}
	low = 0
	high = width
	for low <= high {
		mid := (low + high) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			low = mid + 1
		} else if matrix[row][mid] > target {
			high = mid - 1
		}
	}
	return false
}

func build() [][]int {
	matrix := make([][]int, 6)
	matrix[0] = []int{-8, -7, -5, -3, -3, -1, 1}
	matrix[1] = []int{2, 2, 2, 3, 3, 5, 7}
	matrix[2] = []int{8, 9, 11, 11, 13, 15, 17}
	matrix[3] = []int{18, 18, 18, 20, 20, 20, 21}
	matrix[4] = []int{23, 24, 26, 26, 26, 27, 27}
	matrix[5] = []int{28, 29, 29, 30, 32, 32, 34}
	return matrix
}

func main() {
	matrix := build()
	searchMatrix(matrix, -5)
}
