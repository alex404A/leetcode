package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	m := make(map[string]int)
	return find(s, &m)
}

func find(s string, mPtr *map[string]int) int {
	m := *mPtr
	if value, ok := m[s]; ok {
		return value
	}
	if s[0] == '0' {
		m[s] = 0
		return 0
	}
	if len(s) == 1 {
		m[s] = 1
		return 1
	}
	if len(s) == 2 {
		cnt := 0
		num, _ := strconv.Atoi(s)
		if num <= 26 {
			cnt++
		}
		if s[1] != '0' {
			cnt++
		}
		m[s] = cnt
		return cnt
	}
	cnt := find(s[1:], mPtr)
	num, _ := strconv.Atoi(s[0:2])
	if num <= 26 {
		cnt += find(s[2:], mPtr)
	}
	m[s] = cnt
	return cnt
}

func test(s string, expected int) {
	actual := numDecodings(s)
	if actual != expected {
		fmt.Printf("%s expected %d, actual is %d", s, expected, actual)
	}
}

func main() {
	test("226", 3)
	test("220", 1)
	test("200", 0)
	test("4020", 0)
	test("2020", 1)
}
