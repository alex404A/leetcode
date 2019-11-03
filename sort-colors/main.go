package main

import "fmt"

func sortColors(nums []int) {
	m := map[int]int{
		1: -1,
		2: -1,
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if m[1] > -1 {
				nums[m[1]] = 0
				nums[i] = 1
				if m[2] > -1 {
					nums[m[2]] = 1
					nums[i] = 2
					m[2]++
				}
				m[1]++
			} else if m[2] > -1 {
				nums[m[2]] = 0
				nums[i] = 2
				m[2]++
			}
		} else if nums[i] == 1 {
			if m[2] > -1 {
				nums[m[2]] = 1
				nums[i] = 2
				if m[1] == -1 {
					m[1] = m[2]
				}
				m[2]++
			} else {
				if m[1] == -1 {
					m[1] = i
				}
			}
		} else {
			if m[2] == -1 {
				m[2] = i
			}
		}
	}
	fmt.Println(nums)
}

func main() {
	var nums []int
	nums = []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	nums = []int{2, 2, 2, 1, 2, 0}
	sortColors(nums)
	nums = []int{1, 0, 1, 0, 2, 0, 2, 1}
	sortColors(nums)
	nums = []int{1, 2, 0, 0}
	sortColors(nums)
}
