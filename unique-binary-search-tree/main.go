package main

import "fmt"

func numTrees(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 2,
		3: 5,
	}
	return find(n, &m)
}

func find(n int, mPtr *map[int]int) int {
	m := *mPtr
	val, ok := m[n]
	if ok {
		return val
	} else {
		sum := 0
		for i := 1; i <= n; i++ {
			left := find(i-1, mPtr)
			right := find(n-i, mPtr)
			if left == 0 {
				left = 1
			}
			if right == 0 {
				right = 1
			}
			sum += (left * right)
		}
		m[n] = sum
		return sum
	}
}

func main() {
	result := numTrees(3)
	fmt.Println(result)
}
