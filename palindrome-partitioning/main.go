package main

import "fmt"

type Pair struct {
	solutions [][]string
	ok        bool
}

func partition(s string) [][]string {
	m := iter(s)
	cache := make(map[int]Pair, 0)
	cache[len(s)] = Pair{make([][]string, 0), true}
	solutions, ok := work(s, 0, m, cache)
	if ok {
		return solutions
	} else {
		return make([][]string, 0)
	}
}

func work(s string, cur int, m map[byte][]int, cache map[int]Pair) ([][]string, bool) {
	if pair, ok := cache[cur]; ok {
		return pair.solutions, pair.ok
	}
	b := s[cur]
	indexes := m[b]
	start := biSearch(cur, indexes)
	solutions := make([][]string, 0)
	isAnySolution := false
	for i := start; i < len(indexes); i++ {
		index := indexes[i]
		if isSymmetric(cur, index, indexes) && check(s, cur, index) {
			childSolutions, ok := work(s, index+1, m, cache)
			if ok {
				headSolutions := build(cur, index, s, childSolutions)
				solutions = append(solutions, headSolutions...)
				isAnySolution = true
			}
		}
	}
	cache[cur] = Pair{solutions, isAnySolution}
	return solutions, isAnySolution
}

func build(start int, end int, s string, childSolutions [][]string) [][]string {
	solutions := make([][]string, 0)
	head := s[start : end+1]
	for _, child := range childSolutions {
		solution := make([]string, 0)
		solution = append(solution, head)
		solution = append(solution, child...)
		solutions = append(solutions, solution)
	}
	if len(solutions) == 0 {
		solutions = append(solutions, []string{head})
	}
	return solutions
}

func biSearch(target int, indexes []int) int {
	low := 0
	high := len(indexes) - 1
	for low <= high {
		mid := (low + high) / 2
		if indexes[mid] == target {
			return mid
		} else if indexes[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func iter(s string) map[byte][]int {
	m := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if _, ok := m[b]; !ok {
			m[b] = make([]int, 0)
		}
		m[b] = append(m[b], i)
	}
	return m
}

func isSymmetric(start int, end int, indexes []int) bool {
	innerStartIndex := -1
	innerEndIndex := -1
	for i, index := range indexes {
		if index > start {
			innerStartIndex = i
			break
		}
	}
	for i := len(indexes) - 1; i >= 0; i-- {
		if indexes[i] < end {
			innerEndIndex = i
			break
		}
	}
	if innerStartIndex == -1 || innerEndIndex == -1 {
		return true
	}
	for innerStartIndex <= innerEndIndex {
		if indexes[innerStartIndex]-start != end-indexes[innerEndIndex] {
			return false
		} else {
			innerStartIndex++
			innerEndIndex--
		}
	}
	return true
}

func check(s string, start int, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}

func main() {
	s := partition("aaabaa")
	fmt.Println(s)
}
