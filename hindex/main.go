package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	low := 0
	high := len(citations) - 1
	for high >= low {
		mid := (high + low) / 2
		if citations[mid] >= len(citations)-mid {
			if mid == 0 || citations[mid-1] < len(citations)-mid+1 {
				return len(citations) - mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return 0
}

func test(citations []int, expected int) {
	actual := hIndex(citations)
	if expected != actual {
		fmt.Printf("%v expected %d, actual %d\n", citations, expected, actual)
	}
}

func main() {
	test([]int{}, 0)
	test([]int{1}, 1)
	test([]int{100}, 1)
	test([]int{1, 2, 3, 4, 5}, 3)
	test([]int{3, 0, 6, 1, 5, 3}, 3)
}
