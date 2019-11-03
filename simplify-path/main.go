package main

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	singleSlash = "/"
	singleDot   = "."
	doubleDot   = ".."
	slash       = '/'
	dot         = '.'
)

func simplifyPath(path string) string {
	s := stack{}
	for i := 0; i < len(path); {
		if path[i] == slash {
			i = eatSlash(i, &path)
		} else if path[i] == dot {
			index, name := eatOthers(i, &path)
			i = index
			if name == doubleDot {
				if s.length > 0 {
					s.Pop()
				}
			} else if name != singleDot {
				s.Push(name)
			}
		} else {
			index, name := eatOthers(i, &path)
			i = index
			if s.length > 0 {
				top, _ := s.Get()
				if top == doubleDot {
					s.Pop()
				} else {
					s.Push(name)
				}
			} else {
				s.Push(name)
			}
		}
	}
	buf := bytes.Buffer{}
	for i := 0; i < len(s.stack); i++ {
		buf.WriteString(singleSlash)
		buf.WriteString(s.stack[i])
	}
	result := buf.String()
	if len(result) > 0 {
		return result
	} else {
		return "/"
	}
}

func eatSlash(i int, ptr *string) int {
	path := *ptr
	final := i + 1
	for ; final < len(path) && path[final] == slash; final++ {
	}
	return final
}

func eatOthers(i int, ptr *string) (index int, name string) {
	buf := bytes.Buffer{}
	path := *ptr
	buf.WriteByte(path[i])
	final := i + 1
	for ; final < len(path) && path[final] != slash; final++ {
		buf.WriteByte(path[final])
	}
	return final, buf.String()
}

type stack struct {
	stack  []string
	length int
}

func (s *stack) Push(r string) {
	s.stack = append(s.stack, r)
	s.length++
}

func (s *stack) Pop() (str string, err error) {
	length := len(s.stack)
	if length > 0 {
		result := s.stack[length-1]
		s.stack = s.stack[:length-1]
		s.length--
		return result, nil
	} else {
		return "fail", errors.New("No item in stack")
	}
}

func (s *stack) Get() (str string, err error) {
	length := len(s.stack)
	if length > 0 {
		return s.stack[length-1], nil
	} else {
		return "fail", errors.New("No item in stack")
	}
}

func test(path string, expected string) {
	actual := simplifyPath(path)
	if actual != expected {
		fmt.Printf("%s is expected to be simplified as %s, but actual is %s", path, expected, actual)
	}
}

func main() {
	test("/home/", "/home")
	test("/../", "/")
	test("/home//foo/", "/home/foo")
	test("/a/./b/../../c/", "/c")
	test("/a/../../b/../c//.//", "/c")
	test("/a//b////c/d//././/..", "/a/b/c")
	test("/...", "/...")
	test("/...ab/b/../c", "/...ab/c")
}
