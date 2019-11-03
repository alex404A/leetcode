package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue struct {
	values []*TreeNode
}

func (queue *queue) add(node *TreeNode) {
	queue.values = append(queue.values, node)
}

func (queue *queue) length() int {
	return len(queue.values)
}

func (queue *queue) offer() (*TreeNode, bool) {
	if len(queue.values) > 0 {
		node := queue.values[0]
		queue.values = queue.values[1:]
		return node, true
	} else {
		return nil, false
	}
}

func isSimilarNode(node1 *TreeNode, node2 *TreeNode) bool {
	if node1.Val != node2.Val {
		return false
	}
	if (node1.Left != nil && node2.Left == nil) || (node1.Left == nil && node2.Left != nil) {
		return false
	}
	if (node1.Right != nil && node2.Right == nil) || (node1.Right == nil && node2.Right != nil) {
		return false
	}
	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	q1 := queue{[]*TreeNode{p}}
	q2 := queue{[]*TreeNode{q}}
	for q1.length() > 0 && q2.length() > 0 {
		node1, _ := q1.offer()
		node2, _ := q2.offer()
		isSimilar := isSimilarNode(node1, node2)
		if !isSimilar {
			return false
		} else {
			if node1.Left != nil {
				q1.add(node1.Left)
				q2.add(node2.Left)
			}
			if node1.Right != nil {
				q1.add(node1.Right)
				q2.add(node2.Right)
			}
		}
	}
	return q1.length() == 0 && q2.length() == 0
}

func main() {

}
