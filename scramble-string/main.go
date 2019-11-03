package main

import "fmt"

type block struct {
	before map[byte]int
	after  map[byte]int
}

func isScramble(s1 string, s2 string) bool {
	if !isSame(s1, s2) {
		return false
	}
	if len(s1) <= 2 {
		return true
	}
	dist := preProcess(s1)
	return checkAll(s1, s2, dist)
}

func isSame(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
		m2[s2[i]]++
	}
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func preProcess(s1 string) map[int]block {
	dist := make(map[int]block)
	for i := 1; i < len(s1); i++ {
		block := block{make(map[byte]int), make(map[byte]int)}
		for j := 0; j < i; j++ {
			b := s1[j]
			block.before[b]++
		}
		for j := i; j < len(s1); j++ {
			b := s1[j]
			block.after[b]++
		}
		dist[i] = block
	}
	return dist
}

func checkAll(s1 string, s2 string, dist map[int]block) bool {
	length := len(s2)
	for i := 1; i < length; i++ {
		b1 := dist[i]
		b2 := dist[length-i]
		var c1 map[byte]int
		var c2 map[byte]int
		var sub1 string
		if i < length-i {
			c1 = b1.before
			c2 = b2.after
			sub1 = s2[:i]
		} else {
			c1 = b1.after
			c2 = b2.before
			sub1 = s2[i:]
		}
		r1 := check(sub1, c1)
		r2 := check(sub1, c2)
		if r1 {
			first := isScramble(s2[:i], s1[:i])
			second := isScramble(s2[i:], s1[i:])
			if first && second {
				return true
			}
		}
		if r2 {
			first := isScramble(s2[:i], s1[length-i:])
			second := isScramble(s2[i:], s1[:length-i])
			if first && second {
				return true
			}
		}
	}
	return false
}

func check(sub string, candidate map[byte]int) bool {
	target := make(map[byte]int)
	for i := 0; i < len(sub); i++ {
		target[sub[i]]++
	}
	if len(target) != len(candidate) {
		return false
	}
	for k1, v1 := range target {
		v2, ok := candidate[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func test(s1 string, s2 string, expected bool) {
	actual := isScramble(s1, s2)
	if actual != expected {
		fmt.Printf("%s %s should be %t\n", s1, s2, actual)
	}
}

func main() {
	test("abcdd", "dbdac", false)
	test("great", "geat", false)
	test("", "", true)
	test("g", "g", true)
	test("great", "rgeat", true)
	test("great", "tgear", true)
	test("great", "tegra", true)
	test("great", "etgar", false)
	test("abcde", "caebd", false)
}
