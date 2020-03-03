package main

import (
	"fmt"
	"strconv"
	"strings"
)

type operator byte

const add = '+'
const minus = '-'
const multiply = '*'
const divide = '/'

type numStack struct {
	list []int
}

func (this *numStack) push(num int) {
	this.list = append(this.list, num)
}

func (this *numStack) pop() (num int, ok bool) {
	if len(this.list) == 0 {
		ok = false
	} else {
		num = this.list[len(this.list)-1]
		this.list = this.list[:len(this.list)-1]
		ok = true
	}
	return
}

type operatorStack struct {
	list []operator
}

func (this *operatorStack) push(op operator) {
	this.list = append(this.list, op)
}

func (this *operatorStack) pop() (op operator, ok bool) {
	if len(this.list) == 0 {
		ok = false
	} else {
		op = this.list[len(this.list)-1]
		this.list = this.list[:len(this.list)-1]
		ok = true
	}
	return
}

func (this *operatorStack) isMultiplyOrDivide() bool {
	if len(this.list) == 0 {
		return false
	} else {
		return this.list[len(this.list)-1] == multiply || this.list[len(this.list)-1] == divide
	}
}

func calculate(s string) int {
	numStack := &numStack{make([]int, 0)}
	operatorStack := &operatorStack{make([]operator, 0)}
	turn := false
	next, num := consumeNum(0, s)
	numStack.push(num)
	i := next
	for i < len(s) {
		if turn {
			next, num := consumeNum(i, s)
			op, _ := operatorStack.pop()
			if op == multiply || op == divide {
				first, _ := numStack.pop()
				result := do(first, num, op)
				numStack.push(result)
			} else if op == add {
				numStack.push(num)
			} else if op == minus {
				numStack.push(0 - num)
			}
			i = next
			turn = false
		} else {
			next, op := consumeOp(i, s)
			i = next
			operatorStack.push(op)
			turn = true
		}
	}
	result := 0
	for _, cur := range numStack.list {
		result += cur
	}
	return result
}

func do(first int, second int, op operator) int {
	if op == multiply {
		return first * second
	} else if op == divide {
		return first / second
	} else if op == add {
		return first + second
	} else if op == minus {
		return first - second
	} else {
		return -1
	}
}

func consumeNum(start int, s string) (next int, num int) {
	isFound := false
	for next = start; next < len(s); next++ {
		if s[next] < 48 || s[next] > 57 {
			if isFound {
				break
			} else {
				continue
			}
		} else {
			isFound = true
		}
	}
	num, _ = strconv.Atoi(strings.TrimSpace(s[start:next]))
	return
}

func consumeOp(start int, s string) (next int, op operator) {
	for next = start; next < len(s); next++ {
		if s[next] == ' ' {
			continue
		}
		if s[next] == add {
			op = add
		} else if s[next] == minus {
			op = minus
		} else if s[next] == multiply {
			op = multiply
		} else if s[next] == divide {
			op = divide
		}
		next++
		return
	}
	return
}

func main() {
	result := calculate("3+5 / 2 ")
	fmt.Println(result)
}
