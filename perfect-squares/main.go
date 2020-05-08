package main

import "fmt"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	mem := make([]int, n+1)
	mem[0] = 0
	return find(mem, n)
}

func find(mem []int, n int) int {
	if n == 0 {
		return 0
	}
	if mem[n] > 0 {
		return mem[n]
	}
	result := MaxInt
	for i := 1; i*i <= n; i++ {
		tmp := find(mem, n-i*i) + 1
		if tmp < result {
			result = tmp
		}
	}
	mem[n] = result
	return result
}

func main() {
	fmt.Println(numSquares(12))

}
