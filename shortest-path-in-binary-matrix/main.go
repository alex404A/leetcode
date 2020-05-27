package main

import "fmt"

type Deque struct {
	list [][]int
}

func (this *Deque) append(i int, j int) {
	this.list = append(this.list, []int{i, j})
}

func (this *Deque) popLeft() (i int, j int, ok bool) {
	if len(this.list) == 0 {
		ok = false
	} else {
		p := this.list[0]
		this.list = this.list[1:]
		i = p[0]
		j = p[1]
		ok = true
	}
	return
}

func shortestPathBinaryMatrix(grid [][]int) int {
	N := len(grid)
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	if grid[0][0] == 1 || grid[N-1][N-1] == 1 {
		return -1
	}
	visited := buildVisited(grid)
	mem := buildMem(grid)
	mem[0][0] = 1
	queue := Deque{make([][]int, 0)}
	queue.append(0, 0)
	for len(queue.list) > 0 {
		i, j, _ := queue.popLeft()
		if visited[i][j] {
			continue
		}
		if i == N-1 && j == N-1 {
			return mem[N-1][N-1]
		}
		visited[i][j] = true
		for _, p := range findNeighbours(i, j, grid) {
			nextI := p[0]
			nextJ := p[1]
			if mem[nextI][nextJ] == 0 {
				mem[nextI][nextJ] = mem[i][j] + 1
				queue.append(nextI, nextJ)
			}
		}
	}
	return -1
}

func findNeighbours(i int, j int, grid [][]int) [][]int {
	dirctions := [][]int{
		[]int{1, 1},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, -1},
		[]int{-1, 1},
		[]int{0, -1},
		[]int{-1, 0},
		[]int{-1, -1},
	}
	N := len(grid)
	result := make([][]int, 0)
	for _, dirction := range dirctions {
		nextI := i + dirction[0]
		nextJ := j + dirction[1]
		if nextI < 0 || nextI >= N || nextJ < 0 || nextJ >= N {
			continue
		}
		if grid[nextI][nextJ] == 0 {
			result = append(result, []int{nextI, nextJ})
		}
	}
	return result
}

func buildVisited(grid [][]int) [][]bool {
	mem := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		mem[i] = make([]bool, len(grid[0]))
	}
	return mem
}

func buildMem(grid [][]int) [][]int {
	mem := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		mem[i] = make([]int, len(grid[0]))
	}
	return mem
}

func main() {
	grid := [][]int{
		[]int{0, 0, 0, 0, 1, 1},
		[]int{0, 1, 0, 0, 1, 0},
		[]int{1, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1},
		[]int{0, 1, 0, 0, 0, 1},
		[]int{0, 0, 1, 0, 0, 0},
	}
	result := shortestPathBinaryMatrix(grid)
	fmt.Println(result)
}
