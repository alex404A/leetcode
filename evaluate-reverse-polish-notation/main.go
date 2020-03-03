package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	s []int
}

func (this *Stack) push(token int) {
	this.s = append(this.s, token)
}

func (this *Stack) pop() (token int, ok bool) {
	if len(this.s) == 0 {
		ok = false
		return
	}
	token = this.s[len(this.s)-1]
	this.s = this.s[0 : len(this.s)-1]
	ok = true
	return
}

type caculate func(first, second int) int

const (
	addition       = "+"
	subtraction    = "-"
	multiplication = "*"
	division       = "/"
)

func getMethod(token string) caculate {
	if token == addition {
		return func(first, second int) int {
			return first + second
		}
	} else if token == subtraction {
		return func(first, second int) int {
			return first - second
		}
	} else if token == multiplication {
		return func(first, second int) int {
			return first * second
		}
	} else if token == division {
		return func(first, second int) int {
			return first / second
		}
	} else {
		return nil
	}
}

func evalRPN(tokens []string) int {
	stack := &Stack{make([]int, 0)}
	for _, token := range tokens {
		c := getMethod(token)
		if c == nil {
			num, _ := strconv.Atoi(token)
			stack.push(num)
		} else {
			second, _ := stack.pop()
			first, _ := stack.pop()
			stack.push(c(first, second))
		}
	}
	result, _ := stack.pop()
	return result
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}
