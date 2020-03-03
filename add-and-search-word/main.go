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

func (this *Node) search(word string) bool {
	val := word[0]
	if val == '.' {
		return this.searchDot(word)
	} else {
		return this.searchCommon(val, word)
	}
}

func (this *Node) searchDot(word string) bool {
	if len(word) == 1 {
		for _, child := range this.children {
			if child.isEnd {
				return true
			}
		}
		return false
	}
	for _, child := range this.children {
		isEnd := child.search(word[1:])
		if isEnd {
			return true
		}
	}
	return false
}

func (this *Node) searchCommon(first byte, word string) bool {
	child, ok := this.children[first]
	if !ok {
		return false
	}
	if len(word) == 1 {
		return child.isEnd
	} else {
		return child.search(word[1:])
	}
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	root := &Node{'0', make(map[byte]*Node), false}
	return Trie{root}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.root.insert(word)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.root.search(word)
}

type WordDictionary struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	trie := ConstructorTrie()
	return WordDictionary{&trie}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	dict := Constructor()
	dict.AddWord("bat")
	fmt.Println(dict.Search("b."))
	//dict.AddWord("bad")
	//dict.AddWord("dad")
	//dict.AddWord("mad")
	//fmt.Println(dict.Search("pad"))
	//fmt.Println(dict.Search("bad"))
	//fmt.Println(dict.Search(".ad"))
	//fmt.Println(dict.Search("b.."))
}
