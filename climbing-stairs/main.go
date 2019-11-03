package main

import "fmt"

func climbStairs(n int) int {
	matrix := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	return find(n, &matrix)
}

func find(n int, ptr *map[int]int) int {
	matrix := *ptr
	if n <= 0 {
		return 0
	} else if matrix[n] > 0 {
		return matrix[n]
	} else {
		result := find(n-1, ptr) + find(n-2, ptr)
		matrix[n] = result
		return result
	}
}

func test(n int, expected int) {
	actual := climbStairs(n)
	if actual != expected {
		fmt.Printf("%d is expected to be climbed by %d ways, but actual ways are %d\n", n, expected, actual)
	}
}

func main() {
	test(5, 8)
	test(6, 13)
}
