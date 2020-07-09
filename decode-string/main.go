package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	list []*Item
}

func (s *Stack) push(x *Item) {
	s.list = append(s.list, x)
}

func (s *Stack) pop() (*Item, bool) {
	if len(s.list) == 0 {
		return nil, false
	} else {
		x := s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
		return x, true
	}
}

type ItemType int8

const number = 0
const text = 1
const leftB = 2
const rightB = 3

type Item struct {
	x interface{}
	t ItemType
}

func decodeString(s string) string {
	stack := &Stack{make([]*Item, 0)}
	start := 0
	for start < len(s) {
		end, item := consume(s, start)
		if item.t == number || item.t == text {
			stack.push(item)
		} else if item.t == rightB {
			unit := stack.decodeUnit()
			stack.push(&Item{unit, text})
		}
		start = end
	}
	return stack.join()
}

func (stack *Stack) join() string {
	result := ""
	for len(stack.list) > 0 {
		item, _ := stack.pop()
		x := item.x.(string)
		result = x + result
	}
	return result
}

func (stack *Stack) decodeUnit() string {
	unit := ""
	for len(stack.list) > 0 {
		item, _ := stack.pop()
		if item.t == text {
			unit = item.x.(string) + unit
		} else if item.t == number {
			repeated := item.x.(int)
			result := ""
			for i := 0; i < repeated; i++ {
				result += unit
			}
			return result
		} else {
			panic(fmt.Errorf("error\n"))
		}
	}
	panic(fmt.Errorf("error\n"))
}

func consume(s string, start int) (end int, item *Item) {
	if isNumber(s[start]) {
		end = start + 1
		for isNumber(s[end]) {
			end++
		}
		x, _ := strconv.Atoi(s[start:end])
		item = &Item{x, number}
	} else if isText(s[start]) {
		end = start + 1
		for end < len(s) && isText(s[end]) {
			end++
		}
		item = &Item{s[start:end], text}
	} else if s[start] == '[' {
		end = start + 1
		item = &Item{"[", leftB}
	} else if s[start] == ']' {
		end = start + 1
		item = &Item{"]", rightB}
	} else {
		panic(fmt.Errorf("index %d can not be parsed\n", start))
	}
	return
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isText(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func test(s string, expected string) {
	actual := decodeString(s)
	if actual != expected {
		fmt.Printf("s %s expected %s actual %s\n", s, expected, actual)
	}
}

func main() {
	test("2[abc]3[cd]ef", "abcabccdcdcdef")
	test("3[a]2[bc]", "aaabcbc")
	test("3[a2[c]]", "accaccacc")
	test("ab4[a2[c]]de", "abaccaccaccaccde")
}
