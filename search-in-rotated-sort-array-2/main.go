package main

import "fmt"

func search(nums []int, target int) bool {
	pivot := findPivot(nums)
	if pivot != -1 {
		if nums[pivot] < target || nums[pivot+1] > target {
			return false
		} else if nums[pivot] >= target && nums[0] <= target {
			nums = nums[0 : pivot+1]
		} else if nums[pivot+1] <= target && target <= nums[len(nums)-1] {
			nums = nums[pivot+1:]
		} else {
			return false
		}
	}
	return biSearch(nums, target)
}

func biSearch(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
	}
	return false
}

func findPivot(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return -1
		}
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if mid == len(nums)-1 {
			return -1
		}
		if nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] == nums[low] && nums[low] == nums[high] {
			firstPivot := findPivot(nums[low : mid+1])
			if firstPivot != -1 {
				return firstPivot
			}
			secondPivot := findPivot(nums[mid : high+1])
			if secondPivot != -1 {
				return secondPivot + mid
			} else {
				return -1
			}
		} else if nums[mid] >= nums[low] {
			low = mid + 1
		} else if nums[mid] <= nums[high] {
			high = mid - 1
		}
	}
	return -1
}

func test(nums []int, target int, expect bool) {
	actual := search(nums, target)
	if actual != expect {
		fmt.Printf("%v and %d unexpected\n", nums, target)
	}
}

func main() {
	test([]int{2, 2, 0, 0, 1, 1}, 0, true)
	test([]int{1, 1, 1, 1}, 1, true)
	test([]int{1, 1, 1, 1, 1}, 1, true)
	test([]int{1, 1, 1, 1}, 3, false)
	test([]int{1, 1, 1, 1, 1}, 3, false)
	test([]int{1, 3, 1, 1, 1}, 3, true)
	test([]int{1, 1, 1, 3, 1}, 2, false)
	test([]int{1, 1, 1, 1, 3, 1}, 1, true)
	test([]int{1, 1, 1, 3, 1}, 3, true)
	test([]int{1, 1, 1, 1, 3, 1}, 3, true)
	test([]int{3, 1, 1}, 3, true)
	test([]int{1}, 0, false)
	test([]int{2, 5, 6, 0, 0, 1, 2}, 0, true)
	test([]int{2, 5, 6, 0, 0, 1, 2}, 3, false)
}
