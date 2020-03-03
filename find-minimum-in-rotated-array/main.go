package main

import "fmt"

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[low] {
			high = mid
			low++
		} else {
			high--
		}
	}
	return nums[low]
}

func test(nums []int, expected int) {
	actual := findMin(nums)
	if actual != expected {
		fmt.Printf("%v expected %d actual %d\n", nums, expected, actual)
	}
}

func main() {
	test([]int{10, 1, 10, 10, 10}, 1)
	test([]int{10, 10, 10, 1, 10}, 1)
	test([]int{3, 3, 1}, 1)
	test([]int{1}, 1)
	test([]int{1, 2}, 1)
	test([]int{2, 1}, 1)
	test([]int{1, 2, 3}, 1)
	test([]int{3, 1, 2}, 1)
	test([]int{3, 2, 2}, 2)
	test([]int{2, 2, 2, 0, 1}, 0)
	test([]int{2, 2, 2, 2, 2, 2}, 2)
}
