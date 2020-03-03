package main

import "fmt"

type Node struct {
	b byte
	m map[byte]*Node
}

func (this *Node) add(sub string) bool {
	if len(sub) == 0 {
		return true
	}
	first := sub[0]
	child, ok := this.m[first]
	if ok {
		return child.add(sub[1:])
	} else {
		child = &Node{first, make(map[byte]*Node)}
		this.m[first] = child
		child.add(sub[1:])
		return false
	}
}

type Trie struct {
	root *Node
}

func (this *Trie) add(sub string) bool {
	return this.root.add(sub)
}

func findRepeatedDnaSequences(s string) []string {
	trie := &Trie{&Node{'0', make(map[byte]*Node)}}
	m := make(map[string]bool)
	for i := 0; i <= len(s)-10; i++ {
		isExist := trie.add(s[i : i+10])
		if isExist {
			m[s[i:i+10]] = true
		}
	}
	result := make([]string, 0)
	for key := range m {
		result = append(result, key)
	}
	return result
}

func main() {
	// actual := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	actual := findRepeatedDnaSequences("AAAAAAAAAAA")
	// actual := findRepeatedDnaSequences("AA")
	fmt.Println(actual)
}
