package main

import "fmt"

type Node struct {
	val      byte
	children map[byte]*Node
	isEnd    bool
}

func (this *Node) insert(word string) {
	val := word[0]
	child, ok := this.children[val]
	if !ok {
		child = &Node{val, make(map[byte]*Node), false}
		this.children[val] = child
	}
	if len(word) == 1 {
		child.isEnd = true
	} else {
		child.insert(word[1:])
	}
}

func (this *Node) search(i int, j int, board [][]byte, mem [][]bool) []string {
	isInPath := mem[i][j]
	if isInPath {
		return make([]string, 0)
	}
	mem[i][j] = true
	result := make([]string, 0)
	val := string([]byte{this.val})
	if this.isEnd {
		result = append(result, val)
	}
	if i > 0 {
		next := board[i-1][j]
		child, ok := this.children[next]
		if ok {
			targets := child.search(i-1, j, board, mem)
			result = supplement(result, val, targets)
		}
	}
	if j > 0 {
		next := board[i][j-1]
		child, ok := this.children[next]
		if ok {
			targets := child.search(i, j-1, board, mem)
			result = supplement(result, val, targets)
		}
	}
	if i < len(board)-1 {
		next := board[i+1][j]
		child, ok := this.children[next]
		if ok {
			targets := child.search(i+1, j, board, mem)
			result = supplement(result, val, targets)
		}
	}
	if j < len(board[0])-1 {
		next := board[i][j+1]
		child, ok := this.children[next]
		if ok {
			targets := child.search(i, j+1, board, mem)
			result = supplement(result, val, targets)
		}
	}
	mem[i][j] = false
	return result
}

func supplement(result []string, val string, postfix []string) []string {
	for _, x := range postfix {
		result = append(result, val+x)
	}
	return result
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor(words []string) *Trie {
	root := &Node{'0', make(map[byte]*Node), false}
	trie := &Trie{root}
	for _, word := range words {
		trie.Insert(word)
	}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.root.insert(word)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(i int, j int, board [][]byte, mem [][]bool) []string {
	child, ok := this.root.children[board[i][j]]
	if !ok {
		return make([]string, 0)
	}
	return child.search(i, j, board, mem)
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return make([]string, 0)
	}
	trie := Constructor(words)
	results := make(map[string]bool)
	mem := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		mem[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			targets := trie.Search(i, j, board, mem)
			if len(targets) > 0 {
				for _, c := range targets {
					results[c] = true
				}
			}
		}
	}
	s := make([]string, 0)
	for c := range results {
		s = append(s, c)
	}
	return s
}

func main() {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain", "aaa", "iii"}
	// board := [][]byte{
	// 	[]byte{'a', 'b'},
	// }
	// words := []string{"ab"}
	result := findWords(board, words)
	fmt.Println(result)
}
