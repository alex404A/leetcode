package main

import "fmt"

type Node struct {
	row int
	col int
}

func minimumTotal(triangle [][]int) int {
	if triangle == nil || len(triangle) == 0 {
		return 0
	}
	m := make(map[Node]int)
	return check(Node{0, 0}, triangle, m)
}

func check(root Node, triangle [][]int, m map[Node]int) int {
	val, ok := m[root]
	if ok {
		return val
	}
	if root.row == len(triangle)-1 {
		m[root] = triangle[root.row][root.col]
		return triangle[root.row][root.col]
	}
	left := Node{root.row + 1, root.col}
	right := Node{root.row + 1, root.col + 1}
	leftSum := check(left, triangle, m)
	rightSum := check(right, triangle, m)
	sum := min(leftSum, rightSum) + triangle[root.row][root.col]
	m[root] = sum
	return sum
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	triangle := make([][]int, 0)
	// triangle = append(triangle, []int{2})
	// triangle = append(triangle, []int{3, 4})
	// triangle = append(triangle, []int{6, 5, 7})
	// triangle = append(triangle, []int{4, 1, 8, 3})
	total := minimumTotal(triangle)
	fmt.Println(total)
}
