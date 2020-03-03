package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type Bucket struct {
	min int
	max int
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min, max := findEnds(nums)
	gap := calculateGap(min, max, nums)
	buckets := buildBuckets(min, max, nums, gap)
	return find(gap, buckets)
}

func findEnds(nums []int) (min int, max int) {
	min = MaxInt
	max = MinInt
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

func calculateGap(min int, max int, nums []int) int {
	a := (max - min) / (len(nums) - 1)
	b := (max - min) % (len(nums) - 1)
	if b > 0 {
		return a + 1
	} else {
		return a
	}
}

func buildBuckets(min int, max int, nums []int, gap int) []*Bucket {
	buckets := make([]*Bucket, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		buckets[i] = &Bucket{MaxInt, MinInt}
	}
	for _, num := range nums {
		if num == min {
			buckets[0].min = num
			if buckets[0].max < num {
				buckets[0].max = num
			}
			continue
		}
		if num == max {
			buckets[len(buckets)-1].max = num
			if buckets[len(buckets)-1].min > num {
				buckets[len(buckets)-1].min = num
			}
			continue
		}
		index := (num - min) / gap
		if num < buckets[index].min {
			buckets[index].min = num
		}
		if num > buckets[index].max {
			buckets[index].max = num
		}
	}
	return buckets
}

func find(gap int, buckets []*Bucket) int {
	maxGap := gap
	previous := buckets[0].min
	for _, bucket := range buckets {
		if bucket.min == MaxInt && bucket.max == MinInt {
			continue
		}
		if bucket.min-previous > maxGap {
			maxGap = bucket.min - previous
		}
		previous = bucket.max
	}
	return maxGap
}

func test(nums []int, expected int) {
	actual := maximumGap(nums)
	if actual != expected {
		fmt.Printf("%v expected %d, actual %d", nums, expected, actual)
	}
}

func main() {
	test([]int{1, 10000}, 9999)
	test([]int{3, 6, 9, 1}, 3)
}
