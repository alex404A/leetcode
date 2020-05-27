package main

import (
  "fmt"
)

func maxCoins(nums []int) int {
  addedNums := initCoins(nums)
  mem := initMem(len(addedNums))
  return burst(addedNums, mem, 0, len(addedNums) - 1)
}

func initCoins(nums []int) []int {
  other := make([]int, len(nums) + 2)
  other[0] = 1
  other[len(other) - 1] = 1
  for i := 1; i < len(other) -1; i++ {
    other[i] = nums[i-1]
  }
  return other
}

func initMem(size int) [][]int {
  mem := make([][]int, size)
  for i:= 0; i < size; i++ {
    mem[i] = make([]int, size)
  }
  return mem
}

func burst(nums []int, mem [][]int, left int, right int) int {
  if left + 1 == right {
    return 0
  }
  if mem[left][right] > 0 {
    return mem[left][right]
  }
  max := 0
  for i := left + 1; i < right; i++ {
    max = getMax(max, nums[left] * nums[i] * nums[right] + burst(nums, mem, left, i) + burst(nums, mem, i, right))
  }
  mem[left][right] = max
  return max
}

func getMax(a int, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func main() {
  nums := []int{3,1,5,8}
  result := maxCoins(nums)
  fmt.Println(result)
}
