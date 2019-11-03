package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	if newInterval[0] > intervals[len(intervals)-1][1] {
		return append(intervals, [][]int{newInterval}...)
	}

	var isStartMerged bool
	var isEndMerged bool
	start := -1
	end := -1
	flag := -1

	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		if newInterval[0] > flag && newInterval[0] <= interval[1] {
			start = i
			if newInterval[0] >= interval[0] {
				isStartMerged = true
			}
		}
		if newInterval[1] > flag && newInterval[1] <= interval[1] {
			end = i
			if newInterval[1] >= interval[0] {
				isEndMerged = true
			}
			break
		}
		flag = interval[1]
	}

	mergedInterval := make([]int, 2)
	if isStartMerged {
		mergedInterval[0] = intervals[start][0]
	} else {
		mergedInterval[0] = newInterval[0]
	}
	if isEndMerged {
		mergedInterval[1] = intervals[end][1]
	} else {
		mergedInterval[1] = newInterval[1]
		end--
	}

	return rebuild(&intervals, mergedInterval, start, end)

}

func rebuild(ptr *[][]int, mergedInterval []int, start int, end int) [][]int {
	intervals := *ptr
	var first [][]int
	var last [][]int

	if start >= 0 && start <= len(intervals) {
		first = make([][]int, start)
		copy(first, intervals[:start])
	} else {
		first = make([][]int, 0, 0)
	}

	if end >= -1 && end < len(intervals) {
		last = make([][]int, len(intervals)-end-1)
		copy(last, intervals[end+1:])
	} else {
		last = make([][]int, 0)
	}

	final := append(first, [][]int{mergedInterval}...)
	final = append(final, last...)
	return final
}

func test(intervals [][]int, newInterval []int, expected [][]int) {
	result := insert(intervals, newInterval)
	if len(result) != len(expected) {
		fmt.Printf("newInterval %v, expected is %v, and actual is %v\n", newInterval, expected, result)
		return
	}
	for i := 0; i < len(expected); i++ {
		e := expected[i]
		a := result[i]
		if a[0] != e[0] || a[1] != e[1] {
			fmt.Printf("newInterval %v, expected is %v, and actual is %v\n", newInterval, expected, result)
			break
		}
	}
}

func main() {
	intervals := [][]int{
		[]int{3, 5},
		[]int{12, 12},
	}

	var newInterval []int
	// newInterval = []int{6, 11}
	// test(intervals, newInterval, [][]int{[]int{2, 3}, []int{4, 5}, []int{6, 11}, []int{12, 16}})
	// newInterval = []int{4, 8}
	// test(intervals, newInterval, [][]int{[]int{2, 3}, []int{4, 10}, []int{12, 16}})
	newInterval = []int{6, 6}
	test(intervals, newInterval, [][]int{[]int{3, 5}, []int{6, 6}, []int{12, 12}})
	// newInterval = []int{0, 20}
	// test(intervals, newInterval, [][]int{[]int{0, 20}})
	// newInterval = []int{0, 1}
	// test(intervals, newInterval, [][]int{[]int{0, 1}, []int{2, 3}, []int{4, 5}, []int{8, 10}, []int{12, 16}})
	// newInterval = []int{20, 21}
	// test(intervals, newInterval, [][]int{[]int{2, 3}, []int{4, 5}, []int{8, 10}, []int{12, 16}, []int{20, 21}})
}
