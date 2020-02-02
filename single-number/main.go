package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return nums[0]
	} else {
		m := make(map[int]int)
		for _, num := range nums {
			_, ok := m[num]
			if !ok {
				m[num] = 1
			} else if m[num] == 1 {
				delete(m, num)
			}
		}
		for num := range m {
			return num
		}
		return -1
	}
}

func main() {
	num := singleNumber([]int{4, 1, 3, 1, 3})
	fmt.Println(num)
}
