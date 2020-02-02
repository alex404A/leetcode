package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return make([][]int, 0)
	}
	results := make([][]int, numRows)
	results[0] = []int{1}
	for i := 1; i <= numRows-1; i++ {
		results[i] = build(results[i-1])
	}
	return results
}

func build(row []int) []int {
	next := make([]int, len(row)+1)
	next[0] = row[0]
	for i := 1; i < len(next)-1; i++ {
		next[i] = row[i-1] + row[i]
	}
	next[len(next)-1] = row[len(row)-1]
	return next
}

func main() {
	result := generate(5)
	fmt.Println(result)
}
