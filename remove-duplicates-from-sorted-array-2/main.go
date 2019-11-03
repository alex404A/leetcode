package main

import "fmt"

type info struct {
	last int
	nums []int
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	info := info{len(nums), nums}
	firstNum := nums[0]
	cnt := 1
	for i := 1; i < info.last; i++ {
		num := nums[i]
		if num == firstNum {
			if cnt < 2 {
				cnt++
			} else {
				repeatedLater := info.eat(i, firstNum)
				info.swap(i, i+repeatedLater-1)
				i--
				cnt = 0
			}
		} else {
			firstNum = num
			cnt = 1
		}
	}
	return info.last
}

func (info *info) eat(start int, num int) int {
	result := 0
	for i := start; i < info.last; i++ {
		if num == info.nums[i] {
			result++
		} else {
			break
		}
	}
	return result
}

func (info *info) swap(start int, end int) {
	num := info.nums[start]
	interval := end - start + 1
	isFinished := false
	for i := 0; isFinished == false; i++ {
		for j := start; j <= end; j++ {
			if j+interval < info.last {
				tmp := info.nums[j+interval]
				info.nums[j+interval] = num
				info.nums[j] = tmp
			} else {
				isFinished = true
				info.last = j
				break
			}
			start += interval
			end += interval
		}
	}
}

func test(nums []int, expected int) {
	actual := removeDuplicates(nums)
	fmt.Printf("nums %v\n", nums)
	if actual != expected {
		fmt.Printf("newInterval %v, expected is %d, and actual is %d\n", nums, expected, actual)
	}
}

func main() {
	test([]int{1, 1, 1, 2, 2, 2, 3, 3}, 6)
	test([]int{1, 1, 1, 2, 2, 3}, 5)
	test([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7)
	test([]int{0, 1}, 2)
	test([]int{1, 1}, 2)
	test([]int{1, 1, 1}, 2)
}
