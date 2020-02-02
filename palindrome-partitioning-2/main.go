package main

import "fmt"

type Pair struct {
	ok       bool
	minSplit int
}

func minCut(s string) int {
	m := iter(s)
	cache := make(map[int]Pair, 0)
	cache[len(s)] = Pair{true, 0}
	ok, minSplit := work(s, 0, m, cache)
	if ok {
		return minSplit
	} else {
		return len(s) - 1
	}
}

func work(s string, cur int, m map[byte][]int, cache map[int]Pair) (bool, int) {
	if pair, ok := cache[cur]; ok {
		return pair.ok, pair.minSplit
	}
	minSplit := len(s) - cur - 1
	b := s[cur]
	indexes := m[b]
	start := biSearch(cur, indexes)
	isAnySolution := false
	for i := start; i < len(indexes); i++ {
		index := indexes[i]
		if isSymmetric(cur, index, indexes) && check(s, cur, index) {
			ok, childMinSplit := work(s, index+1, m, cache)
			if index+1 >= len(s) {
				childMinSplit = -1
			}
			if ok {
				minSplit = getMin(minSplit, childMinSplit+1)
				isAnySolution = true
			}
		}
	}
	cache[cur] = Pair{isAnySolution, minSplit}
	return isAnySolution, minSplit
}

func getMin(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
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
	var s int
	s = minCut("aab")
	fmt.Println(s)
	s = minCut("aabbaa")
	fmt.Println(s)
	s = minCut("aabbaaccdffdee")
	fmt.Println(s)
	s = minCut("ababababababababababababcbabababababababababababa")
	fmt.Println(s)
}
