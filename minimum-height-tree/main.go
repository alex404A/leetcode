package main

import (
	"fmt"
	"sort"
)

type Result struct {
	nodes  []int
	height int
}

type Node struct {
	layer int
	val   int
}

type Queue struct {
	list []*Node
}

func (q *Queue) add(node *Node) {
	q.list = append(q.list, node)
}

func (q *Queue) remove() (node *Node, ok bool) {
	if len(q.list) == 0 {
		ok = false
		node = nil
	} else {
		ok = true
		node = q.list[0]
		q.list = q.list[1:]
	}
	return
}

func (q *Queue) size() int {
	return len(q.list)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 2 {
		return []int{0, 1}
	}
	paths := initPath(n, edges)
	nodes := sortNode(paths)
	result := &Result{[]int{0}, n}
	for _, node := range nodes {
		if len(paths[node[0]]) == 1 {
			return result.nodes
		}
		height := calcheight(node[0], paths)
		if height < result.height {
			result.height = height
			result.nodes = []int{node[0]}
		} else if height == result.height {
			result.nodes = append(result.nodes, node[0])
		}
	}
	return result.nodes
}

func calcheight(n int, paths [][]int) int {
	queue := Queue{make([]*Node, 0)}
	node := &Node{
		layer: 0,
		val:   n,
	}
	visited := make([]int, len(paths))
	queue.add(node)
	for queue.size() > 0 {
		node, _ = queue.remove()
		visited[node.val] = 1
		connection := paths[node.val]
		for _, neighbor := range connection {
			if visited[neighbor] == 0 {
				next := &Node{
					layer: node.layer + 1,
					val:   neighbor,
				}
				queue.add(next)
			}
		}
	}
	return node.layer
}

func sortNode(paths [][]int) [][]int {
	nodes := make([][]int, len(paths))
	for i := 0; i < len(paths); i++ {
		nodes[i] = []int{i, len(paths[i])}
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i][1] > nodes[j][1]
	})
	return nodes
}

func initPath(n int, edges [][]int) [][]int {
	paths := make([][]int, n)
	for i := 0; i < n; i++ {
		paths[i] = []int{}
	}
	for _, edge := range edges {
		paths[edge[0]] = append(paths[edge[0]], edge[1])
		paths[edge[1]] = append(paths[edge[1]], edge[0])
	}
	return paths
}

func main() {
	// n := 7
	// edges := [][]int{
	// 	[]int{0, 1},
	// 	[]int{2, 1},
	// 	[]int{3, 1},
	// 	[]int{3, 4},
	// 	[]int{5, 4},
	// 	[]int{6, 4},
	// }
	n := 6
	edges := [][]int{
		[]int{0, 3}, []int{1, 3}, []int{2, 3}, []int{4, 3}, []int{5, 4},
	}

	nodes := findMinHeightTrees(n, edges)
	fmt.Printf("%v\n", nodes)
}
