package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	arr := initArr(n)
	tl := Vertex{0, 0}
	br := Vertex{n - 1, n - 1}
	num := 1
	for ; tl.x <= br.x && tl.y <= br.y; tl.x, br.x, tl.y, br.y = tl.x+1, br.x-1, tl.y+1, br.y-1 {
		if tl.x == br.x && tl.y == br.y {
			arr[tl.x][tl.y] = num
			num++
			break
		}
		for i := tl.y; i < br.y; i++ {
			arr[tl.x][i] = num
			num++
		}
		for i := tl.x; i < br.x; i++ {
			arr[i][br.y] = num
			num++
		}
		for i := br.y; i > tl.y; i-- {
			arr[br.x][i] = num
			num++
		}
		for i := br.x; i > tl.x; i-- {
			arr[i][tl.y] = num
			num++
		}
	}
	return arr
}

func initArr(n int) [][]int {
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	return a
}

type Vertex struct {
	x int
	y int
}

func main() {
	n := 1
	arr := generateMatrix(n)
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}

}
