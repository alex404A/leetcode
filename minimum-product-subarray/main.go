package main

import "fmt"

type Tuple struct {
	index int
	value int
}

func buildIsolationPoints(nums []int) (negativePoints []*Tuple, zeroPoints []*Tuple) {
	negativePoints = make([]*Tuple, 0)
	zeroPoints = make([]*Tuple, 0)
	for i, num := range nums {
		if num < 0 {
			negativePoints = append(negativePoints, &Tuple{i, num})
		} else if num == 0 {
			zeroPoints = append(zeroPoints, &Tuple{i, num})
		}
	}
	return
}

func product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}

func calculate(start int, end int, nums []int, negativePoints []*Tuple) int {
	if len(negativePoints)%2 == 0 {
		return product(nums[start : end+1])
	} else {
		first := product(nums[start:negativePoints[len(negativePoints)-1].index])
		second := product(nums[negativePoints[0].index+1 : end+1])
		if first > second {
			return first
		} else {
			return second
		}
	}
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	negativePoints, zeroPoints := buildIsolationPoints(nums)
	if len(zeroPoints) == 0 {
		return calculate(0, len(nums)-1, nums, negativePoints)
	}
	startNegative := 0
	max := 0
	for i, zeroPoint := range zeroPoints {
		partNegativePoints := make([]*Tuple, 0)
		for j := startNegative; j < len(negativePoints); j++ {
			negativePoint := negativePoints[j]
			if negativePoint.index < zeroPoint.index {
				partNegativePoints = append(partNegativePoints, negativePoint)
				startNegative++
			} else {
				break
			}
		}
		var start, end int
		if i == 0 {
			start = 0
			end = zeroPoint.index - 1
		} else {
			start = zeroPoints[i-1].index + 1
			end = zeroPoint.index - 1
		}
		tmp := calculate(start, end, nums, partNegativePoints)
		if tmp > max {
			max = tmp
		}
	}
	partNegativePoints := negativePoints[startNegative:]
	tmp := calculate(zeroPoints[len(zeroPoints)-1].index+1, len(nums)-1, nums, partNegativePoints)
	if tmp > max {
		max = tmp
	}
	return max
}

func test(nums []int, expected int) {
	actual := maxProduct(nums)
	if expected != actual {
		fmt.Printf("%v expected %d, actual %d", nums, expected, actual)
	}
}

func main() {
	test([]int{2, 3, -2, 4}, 6)
	test([]int{-2, 0, -1}, 0)
	test([]int{-1, 1, -2, 1, 2, -2, 3, 2, 0, 2, -1, 2}, 48)
}
