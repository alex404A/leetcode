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

func (this *Node) search(word string, isPrefixSearch bool) bool {
	val := word[0]
	child, ok := this.children[val]
	if !ok {
		return false
	}
	if len(word) == 1 {
		if isPrefixSearch {
			return true
		} else {
			return child.isEnd
		}
	} else {
		return child.search(word[1:], isPrefixSearch)
	}
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := &Node{'0', make(map[byte]*Node), false}
	return Trie{root}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.root.insert(word)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.root.search(word, false)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.root.search(prefix, true)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
