package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node  *TreeNode
	layer int
}

type Queue struct {
	list   []*Node
	length int
}

func (queue *Queue) offer(node *Node) {
	queue.list = append(queue.list, node)
	queue.length++
}

func (queue *Queue) poll() (*Node, bool) {
	if len(queue.list) == 0 {
		return nil, false
	} else {
		node := queue.list[0]
		queue.list = queue.list[1:]
		queue.length--
		return node, true
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := int(^uint(0) >> 1)
	queue := Queue{make([]*Node, 0), 0}
	node := &Node{root, 1}
	queue.offer(node)
	for queue.length > 0 {
		node, _ := queue.poll()
		if node.layer > result {
			continue
		}
		if node.node.Left == nil && node.node.Right == nil {
			if node.layer < result {
				result = node.layer
			}
			continue
		}
		if node.node.Left != nil {
			left := &Node{node.node.Left, node.layer + 1}
			queue.offer(left)
		}
		if node.node.Right != nil {
			right := &Node{node.node.Right, node.layer + 1}
			queue.offer(right)
		}
	}
	return result
}

func main() {

}
