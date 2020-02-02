package main

import (
	"fmt"
	"sort"
)

type Trie struct {
	root *Node
}

type Node struct {
	s        string
	children []*Node
	end      bool
}

func (trie *Trie) find(s string, index int) []int {
	root := trie.root
	result := make([]int, 0)
	for i := index; i < len(s); i++ {
		b := s[i : i+1]
		child, ok := biSearch(b, root)
		if !ok {
			break
		}
		if child.end {
			result = append(result, i)
		}
		root = child
	}
	return result
}

func biSearch(target string, node *Node) (child *Node, ok bool) {
	low := 0
	high := len(node.children) - 1
	for low <= high {
		mid := (low + high) / 2
		childNode := node.children[mid]
		if childNode.s == target {
			child = childNode
			ok = true
			return
		}
		if childNode.s < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	child = nil
	ok = false
	return
}

type Mem struct {
	mem     map[int][]string
	answers []string
}

func wordBreak(s string, wordDict []string) []string {
	root := &Node{"", make([]*Node, 0), false}
	trie := &Trie{root}
	buildTree(trie.root, wordDict)
	mem := &Mem{make(map[int][]string), make([]string, 0)}
	return check(s, 0, trie, mem)
}

func check(s string, index int, trie *Trie, mem *Mem) []string {
	if value, ok := mem.mem[index]; ok {
		return value
	}
	candidates := trie.find(s, index)
	collection := make([]string, 0)
	for _, candidate := range candidates {
		if candidate == len(s)-1 {
			collection = append(collection, s[index:candidate+1])
		} else {
			childCollection := check(s, candidate+1, trie, mem)
			if len(childCollection) > 0 {
				collection = append(collection, cmb(s[index:candidate+1], childCollection)...)
			}
		}
	}
	mem.mem[index] = collection
	return collection
}

func cmb(before string, after []string) []string {
	dst := make([]string, len(after))
	i := 0
	for ; i < len(after); i++ {
		dst[i] = before + " " + after[i]
	}
	return dst
}

type TmpNode struct {
	start    string
	children []string
}

func buildTree(root *Node, wordDict []string) {
	m := make(map[string][]string)
	for _, word := range wordDict {
		first := word[0:1]
		if _, ok := m[first]; !ok {
			m[first] = make([]string, 0)
		}
		if len(word) > 1 {
			m[first] = append(m[first], word[1:])
		} else {
			m[first] = append(m[first], "")
		}
	}
	tmpNodes := make([]*TmpNode, 0)
	for k, v := range m {
		tmpNodes = append(tmpNodes, &TmpNode{k, v})
	}
	sort.SliceStable(tmpNodes, func(i, j int) bool {
		return tmpNodes[i].start < tmpNodes[j].start
	})
	for _, tmpNode := range tmpNodes {
		node := &Node{tmpNode.start, make([]*Node, 0), false}
		root.children = append(root.children, node)
		childDict := make([]string, 0)
		for _, remain := range tmpNode.children {
			if remain == "" {
				node.end = true
			} else {
				childDict = append(childDict, remain)
			}
		}
		buildTree(node, childDict)
	}
}

func test(s string, wordDict []string) {
	actual := wordBreak(s, wordDict)
	fmt.Printf("%s, %v as model\n", s, wordDict)
	for _, value := range actual {
		fmt.Printf("%s\n", value)
	}
}

func main() {
	test("leetcode", []string{"leet", "code"})
	test("applepenapple", []string{"apple", "pen"})
	test("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	test("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	test("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
	// test("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"})
}
