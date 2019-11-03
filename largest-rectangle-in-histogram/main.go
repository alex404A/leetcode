package main

import "fmt"

func largestRectangleArea(heights []int) int {
	// key is height, value is index
	max := 0
	m := make(map[int]int)
	for i := 0; i < len(heights); i++ {
		height := heights[i]
		for key, value := range m {
			if key > height {
				if max < value {
					max = value
				}
				delete(m, key)
			} else {
				m[key] += key
			}
		}
		_, ok := m[height]
		if !ok {
			m[height] = traceback(&heights, i)
		}
	}
	for _, value := range m {
		if max < value {
			max = value
		}
	}
	return max
}

func traceback(ptr *[]int, index int) int {
	heights := *ptr
	area := 0
	height := heights[index]
	for i := index; i >= 0; i-- {
		if heights[i] >= height {
			area += height
		} else {
			break
		}
	}
	return area
}

func main() {
	var heights []int
	var result int
	// heights = []int{2, 1, 5, 6, 2, 3}
	// result = largestRectangleArea(heights)
	// fmt.Println(result)
	heights = []int{1}
	result = largestRectangleArea(heights)
	fmt.Println(result)
}
