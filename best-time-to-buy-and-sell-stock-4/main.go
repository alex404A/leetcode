package main

import (
	"fmt"
	"sort"
	"strconv"
)

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 || k == 0 {
		return 0
	}
	increase, _ := build(prices)
	return solve(increase, 0, len(increase)-1, k, make(map[string]int), make(map[string]int))
}

func build(prices []int) (increase [][]int, decrease [][]int) {
	increase = make([][]int, 0)
	decrease = make([][]int, 0)
	isIncreasing := prices[1] >= prices[0]
	start := 0
	for i := 2; i < len(prices); i++ {
		if isIncreasing {
			if prices[i] < prices[i-1] {
				increase = append(increase, []int{prices[start], prices[i-1]})
				start = i - 1
				isIncreasing = false
			}
		} else {
			if prices[i] >= prices[i-1] {
				decrease = append(decrease, []int{prices[start], prices[i-1]})
				start = i - 1
				isIncreasing = true
			}
		}
	}
	if isIncreasing {
		increase = append(increase, []int{prices[start], prices[len(prices)-1]})
	} else {
		decrease = append(decrease, []int{prices[start], prices[len(prices)-1]})
	}
	return
}

func solve(increase [][]int, start int, end int, transactions int, mem map[string]int, one map[string]int) int {
	// fmt.Printf("start: %d, end: %d, transactions: %d\n", start, end, transactions)
	if transactions == 0 || start > end {
		return 0
	}
	key := buildKey(start, end, transactions)
	if _, ok := mem[key]; ok {
		return mem[key]
	}
	if transactions == 1 || start >= end {
		result := solveOneTransaction(increase, start, end, one)
		mem[key] = result
		return result
	}
	result := 0
	for i := 0; i <= transactions; i++ {
		first := i
		second := transactions - i
		for j := start; j < end; j++ {
			r1 := solve(increase, start, j, first, mem, one)
			r2 := solve(increase, j+1, end, second, mem, one)
			if r1+r2 > result {
				result = r1 + r2
			}
		}
		// fmt.Printf("start: %d, end: %d, i: %d\n", start, end, i)
	}
	mem[key] = result
	return result
}

func solveOneTransaction(increase [][]int, start int, end int, mem map[string]int) int {
	key := buildOneKey(start, end)
	if value, ok := mem[key]; ok {
		return value
	}
	m := make(map[int]bool)
	m[0] = true
	for i := start; i <= end; i++ {
		for j := i; j <= end; j++ {
			if increase[j][1]-increase[i][0] > 0 {
				m[increase[j][1]-increase[i][0]] = true
			}
		}
	}
	result := make([]int, 0)
	for key := range m {
		result = append(result, key)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	if len(result) == 0 {
		mem[key] = 0
		return 0
	} else {
		mem[key] = result[0]
		return result[0]
	}
}

func buildKey(start, end, transactions int) string {
	return strconv.Itoa(start) + "/" + strconv.Itoa(end) + "/" + strconv.Itoa(transactions)
}

func buildOneKey(start, end int) string {
	return strconv.Itoa(start) + "/" + strconv.Itoa(end)
}

func test(prices []int, k int, expected int) {
	actual := maxProfit(k, prices)
	if actual != expected {
		fmt.Printf("%v at most %d transactions, expected %d, actual %d\n", prices, k, expected, actual)
	}
}
func main() {
	test([]int{3, 2, 6, 5, 0, 3}, 2, 7)
	test([]int{2, 4, 1}, 2, 2)
	test([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 2, 13)
	test([]int{70, 4, 83, 56, 94, 72, 78, 43, 2, 86, 65, 100, 94, 56, 41, 66, 3, 33, 10, 3, 45, 94, 15, 12, 78, 60, 58, 0, 58, 15, 21, 7, 11, 41, 12, 96, 83, 77, 47, 62, 27, 19, 40, 63, 30, 4, 77, 52, 17, 57, 21, 66, 63, 29, 51, 40, 37, 6, 44, 42, 92, 16, 64, 33, 31, 51, 36, 0, 29, 95, 92, 35, 66, 91, 19, 21, 100, 95, 40, 61, 15, 83, 31, 55, 59, 84, 21, 99, 45, 64, 90, 25, 40, 6, 41, 5, 25, 52, 59, 61, 51, 37, 92, 90, 20, 20, 96, 66, 79, 28, 83, 60, 91, 30, 52, 55, 1, 99, 8, 68, 14, 84, 59, 5, 34, 93, 25, 10, 93, 21, 35, 66, 88, 20, 97, 25, 63, 80, 20, 86, 33, 53, 43, 86, 53, 55, 61, 77, 9, 2, 56, 78, 43, 19, 68, 69, 49, 1, 6, 5, 82, 46, 24, 33, 85, 24, 56, 51, 45, 100, 94, 26, 15, 33, 35, 59, 25, 65, 32, 26, 93, 73, 0, 40, 92, 56, 76, 18, 2, 45, 64, 66, 64, 39, 77, 1, 55, 90, 10, 27, 85, 40, 95, 78, 39, 40, 62, 30, 12, 57, 84, 95, 86, 57, 41, 52, 77, 17, 9, 15, 33, 17, 68, 63, 59, 40, 5, 63, 30, 86, 57, 5, 55, 47, 0, 92, 95, 100, 25, 79, 84, 93, 83, 93, 18, 20, 32, 63, 65, 56, 68, 7, 31, 100, 88, 93, 11, 43, 20, 13, 54, 34, 29, 90, 50, 24, 13, 44, 89, 57, 65, 95, 58, 32, 67, 38, 2, 41, 4, 63, 56, 88, 39, 57, 10, 1, 97, 98, 25, 45, 96, 35, 22, 0, 37, 74, 98, 14, 37, 77, 54, 40, 17, 9, 28, 83, 13, 92, 3, 8, 60, 52, 64, 8, 87, 77, 96, 70, 61, 3, 96, 83, 56, 5, 99, 81, 94, 3, 38, 91, 55, 83, 15, 30, 39, 54, 79, 55, 86, 85, 32, 27, 20, 74}, 29, 7)
}
