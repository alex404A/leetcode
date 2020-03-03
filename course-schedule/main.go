package main

import "fmt"

type Node struct {
	id  int
	dep []*Node
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := build(numCourses)
	buildDep(prerequisites, m)
	return search(m)
}

func build(numCourses int) map[int]*Node {
	nodes := make(map[int]*Node)
	for i := 0; i < numCourses; i++ {
		nodes[i] = &Node{i, make([]*Node, 0)}
	}
	return nodes
}

func buildDep(prerequisites [][]int, m map[int]*Node) {
	for _, tuple := range prerequisites {
		original, _ := m[tuple[0]]
		dep, _ := m[tuple[1]]
		original.dep = append(original.dep, dep)
	}
}

func search(m map[int]*Node) bool {
	checked := make(map[int]bool)
	inCycle := make(map[int]bool)
	for _, node := range m {
		if _, ok := checked[node.id]; ok {
			continue
		}
		isLoop := isLoop(checked, inCycle, node)
		if isLoop {
			return false
		}
	}
	return true
}

func isLoop(checked map[int]bool, inCycle map[int]bool, node *Node) bool {
	if v, ok := inCycle[node.id]; ok {
		if v {
			return true
		}
	}
	if _, ok := checked[node.id]; ok {
		return false
	}
	inCycle[node.id] = true
	for _, next := range node.dep {
		isLoop := isLoop(checked, inCycle, next)
		if isLoop {
			return true
		}
	}
	inCycle[node.id] = false
	checked[node.id] = true
	return false
}

func main() {
	numCourses := 2
	prerequisites := [][]int{
		[]int{1, 0},
	}
	result := canFinish(numCourses, prerequisites)
	fmt.Println(result)
}
