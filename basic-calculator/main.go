package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	number     int
	isPositive bool
}

type Stack struct {
	nodes []*Node
}

func (this *Stack) push(node *Node) {
	this.nodes = append(this.nodes, node)
}

func (this *Stack) pop() (node *Node, ok bool) {
	if len(this.nodes) == 0 {
		ok = false
		return
	} else {
		node = this.nodes[len(this.nodes)-1]
		this.nodes = this.nodes[:len(this.nodes)-1]
		ok = true
		return
	}
}

func calculate(s string) int {
	currentNode := &Node{0, true}
	stack := &Stack{make([]*Node, 0)}
	for i := 0; i < len(s); {
		b := s[i]
		if b == '(' {
			stack.push(currentNode)
			currentNode = &Node{0, true}
			i++
		} else if b == ')' {
			lastNode, _ := stack.pop()
			if lastNode.isPositive {
				lastNode.number += currentNode.number
			} else {
				lastNode.number -= currentNode.number
			}
			currentNode = lastNode
			i++
		} else if b == '+' || b == '-' {
			if b == '+' {
				currentNode.isPositive = true
			} else {
				currentNode.isPositive = false
			}
			i++
		} else if b == ' ' {
			i++
		} else {
			bytes := []byte{b}
			start := i + 1
			for ; start < len(s); start++ {
				if s[start] >= 48 && s[start] <= 57 {
					bytes = append(bytes, s[start])
				} else {
					break
				}
			}
			number, _ := strconv.Atoi(string(bytes))
			if currentNode.isPositive {
				currentNode.number += number
			} else {
				currentNode.number -= number
			}
			i = start
		}
	}
	return currentNode.number
}

func test(s string, expected int) {
	actual := calculate(s)
	if actual != expected {
		fmt.Printf("%s expected %d actual %d\n", s, expected, actual)
	}
}

func main() {
	test("2147483647", 2147483647)
	test("2 - (5 - 6)", 3)
	test("(1+(4+5+2)-3)+(6+8)", 23)
	test("1+1", 2)
	test("-1+1", 0)
	test("2 - (5 - (3 - 4)   + 6)", -10)
}
