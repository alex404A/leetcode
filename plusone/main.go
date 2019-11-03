package main

import "fmt"

func plusOne(digits []int) []int {
	isCarried := true
	for i := len(digits) - 1; i >= 0 && isCarried; i-- {
		if isCarried {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i]++
				isCarried = false
			}
		} else {
			isCarried = false
		}
	}
	if isCarried {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}

func main() {
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{2, 0}))
	fmt.Println(plusOne([]int{1, 0, 9}))
}
