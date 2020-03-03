package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	i := 1
	for i <= len(nums)-2 {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		} else if nums[i] < nums[i+1] {
			i++
		} else {
			i += 2
		}
	}
	return -1
}

func main() {

}
