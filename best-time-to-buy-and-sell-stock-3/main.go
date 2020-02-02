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
	result := 0
	for i := 0; i < len(prices); i++ {
		left := doSingle(Location{0, i}, prices)
		right := doSingle(Location{i, len(prices) - 1}, prices)
		if result < left+right {
			result = left + right
		}
	}
	return result
}

func doSingle(location Location, prices []int) int {
	if location.end <= location.start {
		return 0
	}
	max := 0
	minPrice := prices[location.start]
	for i := location.start + 1; i <= location.end; i++ {
		price := prices[i]
		sub := price - minPrice
		if max < sub {
			max = sub
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return max
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
	test([]int{3, 3, 5, 0, 0, 3, 1, 4}, 6)
	test([]int{1, 2, 3, 4, 5}, 4)
	test([]int{7, 6, 4, 3, 1}, 0)
}
