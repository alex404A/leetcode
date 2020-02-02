package main

import (
	"fmt"
)

type Location struct {
	start int
	end   int
}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}
	return do(Location{0, len(prices) - 1}, prices)
}

func do(location Location, prices []int) int {
	result := 0
	if location.end-location.start < 3 {
		return check(location, prices)
	}
	left := do(Location{location.start, location.start + 1}, prices)
	right := do(Location{location.start + 1, location.end}, prices)
	if result < left+right {
		result = left + right
	}
	return result
}

func check(location Location, prices []int) int {
	if location.end-location.start < 1 {
		return 0
	} else if location.end-location.start == 1 {
		sub := prices[location.end] - prices[location.start]
		return max(sub, 0)
	} else if location.end-location.start == 2 {
		sub1 := prices[location.end] - prices[location.start]
		sub2 := prices[location.end-1] - prices[location.start]
		sub3 := prices[location.end] - prices[location.start+1]
		sub := max(sub1, sub2, sub3, 0)
		return sub
	} else {
		fmt.Println("error")
		return 0
	}
}

func max(nums ...int) int {
	result := 0
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

func test(prices []int, expected int) {
	actual := maxProfit(prices)
	if expected != actual {
		fmt.Printf("%v\n", prices)
	}
}

func main() {
}
