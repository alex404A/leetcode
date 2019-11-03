package main

import "fmt"

func buildInfo(t string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	return m
}

func minWindow(s string, t string) string {
	m := buildInfo(t)
	begin := 0
	end := 0
	counter := len(t)
	minBegin := 0
	d := len(s) + 1

	for end < len(s) {
		if m[s[end]] > 0 {
			counter--
		}
		m[s[end]]--
		end++
		for counter == 0 {
			if end-begin < d {
				d = end - begin
				minBegin = begin
			}
			m[s[begin]]++
			if m[s[begin]] > 0 {
				counter++
			}
			begin++
		}
	}

	if d == len(s)+1 {
		return ""
	} else {
		return s[minBegin : minBegin+d]
	}
}

func test(s string, t string, expected string) {
	result := minWindow(s, t)
	if result != expected {
		fmt.Printf("search minimum window in %s with %s, expected is %s, and actual is %s\n", s, t, expected, result)
	}
}

func main() {
	test("BABCDEAX", "AD", "DEA")
	test("ADOBECODEBANC", "ABC", "BANC")
	test("ADOBECODEBANC", "BD", "DOB")
	test("ADOBECODEBANC", "DE", "DE")
	test("ADOBECODEBANC", "BXY", "")
	test("a", "b", "")
	test("a", "a", "a")
	test("aa", "aa", "aa")
	test("aaflslflsldkalskaaa", "aaa", "aaa")
	test("aaaaa", "aaaa", "aaaa")
	test("abcabdebac", "cea", "ebac")
}
