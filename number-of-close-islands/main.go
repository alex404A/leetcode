package main

import "fmt"

func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	mem := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		mem[i] = make([]bool, len(grid[0]))
	}
	count := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 0 && !mem[i][j] {
				ok := dfs(i, j, grid, mem)
				if ok {
					count++
				}
			}
		}
	}
	return count
}

func dfs(i int, j int, grid [][]int, mem [][]bool) bool {
	if mem[i][j] {
		return true
	}
	if i == len(mem)-1 || i == 0 {
		return false
	}
	if j == len(mem[0])-1 || j == 0 {
		return false
	}
	mem[i][j] = true
	ok := true
	if grid[i][j+1] == 0 {
		ok = dfs(i, j+1, grid, mem) && ok
	}
	if grid[i][j-1] == 0 {
		ok = dfs(i, j-1, grid, mem) && ok
	}
	if grid[i+1][j] == 0 {
		ok = dfs(i+1, j, grid, mem) && ok
	}
	if grid[i-1][j] == 0 {
		ok = dfs(i-1, j, grid, mem) && ok
	}
	return ok
}

func main() {
	grid := [][]int{
		[]int{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		[]int{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		[]int{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		[]int{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		[]int{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		[]int{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		[]int{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		[]int{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	}
	count := closedIsland(grid)
	fmt.Println(count)
}
