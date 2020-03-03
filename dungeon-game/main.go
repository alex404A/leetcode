package main

import "fmt"

type Point struct {
	x int
	y int
}

func calculateMinimumHP(dungeon [][]int) int {
	m := initial(dungeon)
	result := start(0, 0, dungeon, m)
	return result
}

func start(x int, y int, dungeon [][]int, m [][]int) int {
	if m[x][y] > 0 {
		return m[x][y]
	}
	right := start(x, y+1, dungeon, m)
	right = calculate(right, dungeon[x][y])
	down := start(x+1, y, dungeon, m)
	down = calculate(down, dungeon[x][y])
	if right < down {
		m[x][y] = right
		return right
	} else {
		m[x][y] = down
		return down
	}
}

func initial(dungeon [][]int) [][]int {
	m := make([][]int, len(dungeon))
	for i := 0; i < len(dungeon); i++ {
		m[i] = make([]int, len(dungeon[0]))
	}
	lastRow := len(dungeon) - 1
	lastCol := len(dungeon[0]) - 1
	m[lastRow][lastCol] = calculate(1, dungeon[lastRow][lastCol])
	for i := lastRow - 1; i >= 0; i-- {
		m[i][lastCol] = calculate(m[i+1][lastCol], dungeon[i][lastCol])
	}
	for j := lastCol - 1; j >= 0; j-- {
		m[lastRow][j] = calculate(m[lastRow][j+1], dungeon[lastRow][j])
	}
	return m
}

func calculate(accu int, now int) int {
	if now >= 0 {
		if accu-now <= 0 {
			return 1
		} else {
			return accu - now
		}
	} else {
		return accu - now
	}
}

func main() {
	dungeon := [][]int{
		[]int{-2, -3, 3},
		[]int{-5, -10, 1},
		[]int{10, 30, -5},
	}
	result := calculateMinimumHP(dungeon)
	fmt.Println(result)
}
