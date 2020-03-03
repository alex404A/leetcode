package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	r1 := sub(nums[1:])
	r2 := sub(nums[:len(nums)-1])
	return getMax(r1, r2)
}

func sub(nums []int) int {
	prev1 := 0
	prev2 := 0
	for _, num := range nums {
		tmp := prev1
		prev1 = getMax(prev2+num, prev1)
		prev2 = tmp
	}
	return prev1
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

}
