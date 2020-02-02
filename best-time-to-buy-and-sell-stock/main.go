package main

import "fmt"

func maxProfit(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}
	max := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
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

func main() {
	prices := []int{7, 8}
	result := maxProfit(prices)
	fmt.Println(result)
}
