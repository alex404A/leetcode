package main

import "fmt"

type Node struct {
	id       int
	indegree int
	dep      []*Node
}

type Candidates struct {
	m map[int]*Node
}

func (this *Candidates) put(node *Node) {
	this.m[node.id] = node
}

func (this *Candidates) popNoDegree() (node *Node, ok bool) {
	ok = false
	for id, c := range this.m {
		if c.indegree == 0 {
			node = c
			ok = true
			delete(this.m, id)
			return
		}
	}
	return
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	todo, _ := build(numCourses, prerequisites)
	result := make([]int, 0)
	candidates := &Candidates{make(map[int]*Node)}
	for _, node := range todo {
		if node.indegree == 0 {
			candidates.put(node)
		}
	}
	for len(candidates.m) > 0 {
		c, ok := candidates.popNoDegree()
		if !ok {
			return make([]int, 0)
		}
		result = append(result, c.id)
		for _, dep := range c.dep {
			dep.indegree--
			candidates.put(dep)
		}
	}
	if len(result) != numCourses {
		return make([]int, 0)
	} else {
		return result
	}
}

func build(numCourses int, prerequisites [][]int) (todo map[int]*Node, done map[int]*Node) {
	todo = make(map[int]*Node)
	done = make(map[int]*Node)
	for i := 0; i < numCourses; i++ {
		node := &Node{i, 0, make([]*Node, 0)}
		todo[i] = node
	}
	for _, dep := range prerequisites {
		node, _ := todo[dep[1]]
		next, _ := todo[dep[0]]
		next.indegree++
		node.dep = append(node.dep, next)
	}
	return
}

func main() {
	result := findOrder(4, [][]int{
		[]int{1, 0},
		[]int{2, 0},
		[]int{3, 1},
		[]int{3, 2},
	})
	fmt.Println(result)
}
