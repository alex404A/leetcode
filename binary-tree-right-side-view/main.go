package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	id       string
	treeNode *TreeNode
}

func (node *Node) buildLeft(treeNode *TreeNode) *Node {
	id := node.id + "1"
	return &Node{id, treeNode}
}

func (node *Node) buildRight(treeNode *TreeNode) *Node {
	id := node.id + "2"
	return &Node{id, treeNode}
}

type Queue struct {
	list []*Node
}

func (queue *Queue) offer(node *Node) {
	queue.list = append(queue.list, node)
}

func (queue *Queue) peekFromTail() (node *Node, ok bool) {
	if len(queue.list) == 0 {
		ok = false
		return
	} else {
		ok = true
		node = queue.list[len(queue.list)-1]
		return
	}
}

func (queue *Queue) pollFromTail() (node *Node, ok bool) {
	if len(queue.list) == 0 {
		ok = false
		return
	} else {
		ok = true
		node = queue.list[len(queue.list)-1]
		queue.list = queue.list[:len(queue.list)-1]
		return
	}
}

func (queue *Queue) poll() (node *Node, ok bool) {
	if len(queue.list) == 0 {
		ok = false
		return
	} else {
		ok = true
		node = queue.list[0]
		queue.list = queue.list[1:]
		return
	}
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	first := &Queue{make([]*Node, 0)}
	second := &Queue{make([]*Node, 0)}
	node := &Node{"1", root}
	first.offer(node)

	for len(first.list) > 0 {
		node, _ := first.poll()
		if node.treeNode.Left != nil {
			left := node.buildLeft(node.treeNode.Left)
			first.offer(left)
		}
		if node.treeNode.Right != nil {
			right := node.buildRight(node.treeNode.Right)
			first.offer(right)
		}
		second.offer(node)
	}

	last, _ := second.peekFromTail()
	layer := len(last.id)
	nodes := make([]*Node, 1)
	nodes[0] = last
	for len(second.list) > 0 {
		last, _ = second.pollFromTail()
		if layer != len(last.id) {
			nodes = append(nodes, last)
			layer = len(last.id)
		}
	}

	result := make([]int, len(nodes))
	length := len(nodes) - 1
	for i := len(nodes) - 1; i >= 0; i-- {
		result[length-i] = nodes[i].treeNode.Val
	}

	return result
}

func main() {

}
