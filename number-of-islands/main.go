package main

import (
	"fmt"
	"strconv"
)

const land = '1'
const water = '0'

type Node struct {
	i   int
	j   int
	v   byte
	key string
}

func buildNode(i int, j int, v byte) *Node {
	key := strconv.Itoa(i) + ":" + strconv.Itoa(j)
	return &Node{i, j, v, key}
}

type Set struct {
	m map[string]bool
}

func (set *Set) put(key string) {
	set.m[key] = true
}

func (set *Set) contains(key string) bool {
	_, ok := set.m[key]
	return ok
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	nodes := build(grid)
	return search(nodes)
}

func build(grid [][]byte) [][]*Node {
	nodes := make([][]*Node, len(grid))
	for i := 0; i < len(nodes); i++ {
		nodes[i] = make([]*Node, len(grid[0]))
	}
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes[0]); j++ {
			nodes[i][j] = buildNode(i, j, grid[i][j])
		}
	}
	return nodes
}

func search(nodes [][]*Node) int {
	set := &Set{make(map[string]bool)}
	cnt := 0
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes[0]); j++ {
			node := nodes[i][j]
			if node.v == land && !set.contains(node.key) {
				cnt++
				set.put(node.key)
				explore(node, nodes, set)
			}
		}
	}
	return cnt
}

func explore(node *Node, nodes [][]*Node, set *Set) {
	if node.i > 0 {
		up := nodes[node.i-1][node.j]
		if up.v == land && !set.contains(up.key) {
			set.put(up.key)
			explore(up, nodes, set)
		}
	}
	if node.i < len(nodes)-1 {
		down := nodes[node.i+1][node.j]
		if down.v == land && !set.contains(down.key) {
			set.put(down.key)
			explore(down, nodes, set)
		}
	}
	if node.j > 0 {
		left := nodes[node.i][node.j-1]
		if left.v == land && !set.contains(left.key) {
			set.put(left.key)
			explore(left, nodes, set)
		}
	}
	if node.j < len(nodes[0])-1 {
		right := nodes[node.i][node.j+1]
		if right.v == land && !set.contains(right.key) {
			set.put(right.key)
			explore(right, nodes, set)
		}
	}
}

func main() {
	// grid := [][]byte{
	// 	[]byte{'1', '1', '1', '1', '0'},
	// 	[]byte{'1', '1', '0', '1', '0'},
	// 	[]byte{'1', '1', '0', '0', '0'},
	// 	[]byte{'0', '0', '0', '0', '0'},
	// }
	grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	cnt := numIslands(grid)
	fmt.Println(cnt)
}
