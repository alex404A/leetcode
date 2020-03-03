package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {

	i := 0
	j := 0
	sum := 0
	min := 99999999

	for j < len(nums) {
		sum += nums[j]
		j++

		for sum >= s {
			if j-i < min {
				min = j - i
			}
			sum -= nums[i]
			i++
		}
	}

	if min == 99999999 {
		return 0
	} else {
		return min
	}
}

func test(s int, nums []int, expected int) {
	actual := minSubArrayLen(s, nums)
	if actual != expected {
		fmt.Printf("%v with minimum %d, expected %d, actual %d\n", nums, s, expected, actual)
	}
}

func main() {
	test(11, []int{1, 2, 3, 4, 5}, 3)
	test(7, []int{2, 3, 1, 2, 4, 3}, 2)
	test(7, []int{2, 3, 1, 2, 4, 3, 7}, 1)
	test(7, []int{1, 2, 1, 2}, 0)
	test(7, []int{1, 2, 1, 3}, 4)
}
