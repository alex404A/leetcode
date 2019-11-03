package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	bigger := nums1
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			bigger[k] = nums1[i]
			i--
		} else {
			bigger[k] = nums2[j]
			j--
		}
		k--
	}
	if i >= 0 && &bigger != &nums1 {
		for ; i >= 0; i-- {
			bigger[i] = nums1[i]
		}
	}
	if j >= 0 && &bigger != &nums2 {
		for ; j >= 0; j-- {
			bigger[j] = nums2[j]
		}
	}
}

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, len(nums1)-1, nums2, len(nums2))
	fmt.Println(nums1)
}
