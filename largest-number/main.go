package main

import (
	"fmt"
	"sort"
	"strconv"
)

const seq = "0123456789"

func largestNumber(nums []int) string {
	m := build(nums)
	if len(m) == 1 {
		if _, ok := m['0']; ok {
			return "0"
		}
	}
	sortBucket(m)
	return combine(m)
}

func build(nums []int) map[byte][]string {
	m := make(map[byte][]string)
	for i := range seq {
		m[seq[i]] = make([]string, 0)
	}
	for _, num := range nums {
		s := strconv.Itoa(num)
		m[s[0]] = append(m[s[0]], s)
	}
	for i := range seq {
		arr, _ := m[seq[i]]
		if len(arr) == 0 {
			delete(m, seq[i])
		}
	}
	return m
}

func sortBucket(m map[byte][]string) {
	for i := range seq {
		arr := m[seq[i]]
		if len(arr) == 0 {
			continue
		}
		first := seq[i]
		sort.SliceStable(arr, func(i, j int) bool {
			a := arr[i]
			b := arr[j]
			return do(a, b, first)
		})
		m[seq[i]] = arr
	}
}

func do(a string, b string, first byte) bool {
	length := len(b)
	if len(a) < len(b) {
		length = len(a)
	}
	for i := 1; i < length; i++ {
		if a[i] < b[i] {
			return false
		} else if a[i] > b[i] {
			return true
		}
	}
	if len(a) == len(b) {
		return true
	} else if len(a) > len(b) {
		for i := len(b); i < len(a); i++ {
			if a[i] > first {
				return true
			} else if a[i] < first {
				return false
			}
		}
	} else {
		for i := len(a); i < len(b); i++ {
			if b[i] > first {
				return false
			} else if b[i] < first {
				return true
			}
		}
	}
	x, _ := strconv.Atoi(a + b)
	y, _ := strconv.Atoi(b + a)
	return x >= y
}

func combine(m map[byte][]string) string {
	result := ""
	for i := len(seq) - 1; i >= 0; i-- {
		arr := m[seq[i]]
		for _, str := range arr {
			result += str
		}
	}
	return result
}

func test(nums []int, expected string) {
	actual := largestNumber(nums)
	if actual != expected {
		fmt.Printf("%v expected %s actual %s\n", nums, expected, actual)
	}
}

func main() {
	test([]int{9051, 5526, 2264, 5041, 1630, 5906, 6787, 8382, 4662, 4532, 6804, 4710, 4542, 2116, 7219, 8701, 8308, 957, 8775, 4822, 396, 8995, 8597, 2304, 8902, 830, 8591, 5828, 9642, 7100, 3976, 5565, 5490, 1613, 5731, 8052, 8985, 2623, 6325, 3723, 5224, 8274, 4787, 6310, 3393, 78, 3288, 7584, 7440, 5752, 351, 4555, 7265, 9959, 3866, 9854, 2709, 5817, 7272, 43, 1014, 7527, 3946, 4289, 1272, 5213, 710, 1603, 2436, 8823, 5228, 2581, 771, 3700, 2109, 5638, 3402, 3910, 871, 5441, 6861, 9556, 1089, 4088, 2788, 9632, 6822, 6145, 5137, 236, 683, 2869, 9525, 8161, 8374, 2439, 6028, 7813, 6406, 7519}, "995998549642963295795569525905189958985890288238775871870185978591838283748308830827481618052787813771758475277519744072727265721971071006861683682268046787640663256310614560285906582858175752573156385565552654905441522852245213513750414822478747104662455545424532434289408839763963946391038663723370035134023393328828692788270926232581243924362362304226421162109163016131603127210891014")
	test([]int{0, 0}, "0")
	test([]int{12, 121}, "12121")
	test([]int{121, 12}, "12121")
	test([]int{3, 30, 34, 5, 9}, "9534330")
	test([]int{10, 2}, "210")
	test([]int{332, 3, 341, 333}, "3413333332")
	test([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, "9876543210")
}
