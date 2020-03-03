package main

import "fmt"

type Result struct {
	answers [][]int
	comb    []int
}

func combinationSum3(k int, n int) [][]int {
	result := &Result{make([][]int, 0), make([]int, 0)}
	find(1, k, n, result)
	return result.answers
}

func find(start int, k int, n int, result *Result) {
	if n < 0 {
		return
	}
	if k == 0 {
		if n == 0 {
			result.answers = append(result.answers, copy(result.comb))
		}
		return
	}
	for i := start; i <= 9; i++ {
		result.comb = append(result.comb, i)
		find(i+1, k-1, n-i, result)
		result.comb = result.comb[:len(result.comb)-1]
	}
}

func copy(comb []int) []int {
	result := make([]int, len(comb))
	for i := 0; i < len(comb); i++ {
		result[i] = comb[i]
	}
	return result
}

func main() {
	answers := combinationSum3(3, 9)
	fmt.Println(answers)
}
