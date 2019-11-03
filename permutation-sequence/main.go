package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	s := make([]int, n)
	result := ""
	for i := 0; i < n; i++ {
		s[i] = i + 1
	}
	for i := n - 1; i > 0; i-- {
		f := factorial(i)
		var j int
		if k < f {
			j = 0
		} else if k%f != 0 {
			j = k / f
			k = k % f
		} else {
			j = k/f - 1
			result += strconv.Itoa(s[j])
			s = remove(s, j)
			s = reverse(s)
			break
		}
		result += strconv.Itoa(s[j])
		s = remove(s, j)
	}
	result += join(s)
	fmt.Println(result)
	return result
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = result * i
	}
	return result
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func reverse(slice []int) []int {
	result := make([]int, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		result[len(slice)-i-1] = slice[i]
	}
	return result
}

func join(s []int) string {
	result := ""
	for i := 0; i < len(s); i++ {
		result += strconv.Itoa(s[i])
	}
	return result
}

func main() {
	getPermutation(3, 4)
}
