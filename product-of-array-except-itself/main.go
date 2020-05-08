package main

import "fmt"

func productExceptSelf(nums []int) []int {
	product := 1
	zeroKey := -1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num == 0 {
			if zeroKey == -1 {
				zeroKey = i
			} else {
				return make([]int, len(nums))
			}
		} else {
			product *= num
		}
	}
	results := make([]int, len(nums))
	if zeroKey != -1 {
		results[zeroKey] = product
		return results
	}
	for i := 0; i < len(nums); i++ {
		results[i] = product / nums[i]
	}
	return results
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{1, 2, 0, 4}))
	fmt.Println(productExceptSelf([]int{0, 2, 0, 4}))
}
