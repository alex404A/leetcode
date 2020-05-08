package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node *TreeNode
	id   string
}

func (node *Node) String() string {
	return fmt.Sprintf("%d/%s", node.node.Val, node.id)
}

type Queue struct {
	list []*Node
}

func (queue *Queue) offer(node *Node) {
	queue.list = append(queue.list, node)
}

func (queue *Queue) poll() (node *Node, ok bool) {
	if len(queue.list) == 0 {
		node = nil
		ok = false
	} else {
		node = queue.list[0]
		ok = true
		queue.list = queue.list[1:]
	}
	return
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	arr := []string{}
	if root == nil {
		bytes, _ := json.Marshal(arr)
		return string(bytes)
	}
	queue := Queue{make([]*Node, 0)}
	queue.offer(&Node{root, "1"})
	for len(queue.list) > 0 {
		node, _ := queue.poll()
		arr = append(arr, fmt.Sprintf("%v", node))
		if node.node.Left != nil {
			left := &Node{node.node.Left, node.id + ".1"}
			queue.offer(left)
		}
		if node.node.Right != nil {
			right := &Node{node.node.Right, node.id + ".2"}
			queue.offer(right)
		}
	}
	b, _ := json.Marshal(arr)
	return string(b)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := []string{}
	m := make(map[string]*TreeNode)
	json.Unmarshal([]byte(data), &arr)
	if len(arr) == 0 {
		return nil
	}
	for _, s := range arr {
		slice := strings.Split(s, "/")
		val, _ := strconv.Atoi(slice[0])
		id := slice[1]
		treeNode := &TreeNode{val, nil, nil}
		m[id] = treeNode
		if strings.HasSuffix(id, ".1") {
			parentID := id[0 : len(id)-2]
			parent, _ := m[parentID]
			parent.Left = treeNode
		} else if strings.HasSuffix(id, ".2") {
			parentID := id[0 : len(id)-2]
			parent, _ := m[parentID]
			parent.Right = treeNode
		}
	}
	root, _ := m["1"]
	return root
}

func main() {
	left := &TreeNode{1, nil, nil}
	right := &TreeNode{2, nil, nil}
	root := &TreeNode{0, left, right}
	codec := Constructor()
	data := codec.serialize(root)
	fmt.Println(data)
	other := codec.deserialize(data)
	fmt.Println(other.Val)
}
