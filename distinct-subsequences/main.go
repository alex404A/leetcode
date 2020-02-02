package main

import "fmt"

type Node struct {
	s string
	t string
}

type Location struct {
	pts int
	ptt int
}

func numDistinct(s string, t string) int {
	blueprint := extract(s, t)
	if !verify(blueprint, t) {
		return 0
	}
	if len(t) == 1 {
		return len(blueprint[t[0]])
	}
	mem := make(map[Location]int)
	node := Node{s, t}
	locations := blueprint[node.t[0]]
	next, _ := biSearch(0, locations)
	return check(&node, locations[next], 0, blueprint, mem)
}

func check(node *Node, pts int, ptt int, blueprint map[byte][]int, mem map[Location]int) int {
	if val, ok := mem[Location{pts, ptt}]; ok {
		return val
	}
	sum := 0
	locations := blueprint[node.t[ptt+1]]
	next1, ok1 := biSearch(pts+1, locations)
	if ok1 {
		if len(node.t) == ptt+2 {
			sum += len(locations) - next1
		} else {
			sum += check(node, locations[next1], ptt+1, blueprint, mem)
		}
	}
	locations = blueprint[node.t[ptt]]
	next2, ok2 := biSearch(pts+1, locations)
	if ok2 {
		sum += check(node, locations[next2], ptt, blueprint, mem)
	}
	mem[Location{pts, ptt}] = sum
	return sum
}

func biSearch(target int, locations []int) (int, bool) {
	low := 0
	high := len(locations) - 1
	for low <= high {
		mid := (low + high) / 2
		if mid == 0 {
			if locations[mid] >= target {
				return 0, true
			} else {
				low = mid + 1
				continue
			}
		}
		if locations[mid] >= target && locations[mid-1] < target {
			return mid, true
		}
		if locations[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}

func extract(s string, t string) map[byte][]int {
	m := make(map[byte][]int)
	for i := 0; i < len(t); i++ {
		m[t[i]] = make([]int, 0)
	}
	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if ok {
			m[s[i]] = append(m[s[i]], i)
		}
	}
	return m
}

func verify(m map[byte][]int, t string) bool {
	counter := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		counter[t[i]]++
	}
	for c, _ := range m {
		cnt, _ := counter[c]
		if cnt > len(m[c]) {
			return false
		}
	}
	return true
}

func test(s string, t string, expected int) {
	actual := numDistinct(s, t)
	if actual != expected {
		fmt.Printf("%s %s should be %d, actual is %d\n", s, t, expected, actual)
	}
}

func main() {
	// test("abcd", "ae", 0)
	// test("abcd", "aa", 0)
	// test("rabbbit", "rabbit", 3)
	// test("babgbag", "bag", 5)
	test("bit", "it", 1)

	test("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc", 3)
}
