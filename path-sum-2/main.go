package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Container struct {
	list [][]int
}

func pathSum(root *TreeNode, sum int) [][]int {
	container := Container{make([][]int, 0)}
	check(root, sum, make([]int, 0), &container)
	return container.list
}

func check(root *TreeNode, sum int, path []int, container *Container) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			dst := make([]int, len(path)+1)
			copy(dst, append(path, root.Val))
			container.list = append(container.list, dst)
		} else {
			return
		}
	}
	check(root.Left, sum-root.Val, append(path, root.Val), container)
	check(root.Right, sum-root.Val, append(path, root.Val), container)
}

func build(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 1)
	nodes[0] = &TreeNode{array[0], nil, nil}
	for i := 1; i < len(array); i++ {
		val := array[i]
		if val == -999 {
			continue
		}
		node := &TreeNode{val, nil, nil}
		nodes = append(nodes, node)
		j := i / 2
		k := i % 2
		if k == 1 {
			parent := nodes[j]
			parent.Left = node
		}
		if k == 0 {
			parent := nodes[j-1]
			parent.Right = node
		}
	}
	return nodes[0]
}

func main() {
	nums := []int{-260, -202, -903, -980, -570, -858, 218, 764, -300, 205, -999, -35, -999, -999, -204, 950, -769, 258, -652, 614, -584, 76, 817, -192, -999, -999, -114, 880, -999, -200, 71, 671, 344, 801, 193, -18, 876, -920, -730, 222, 679, -999, -680, -999, -999, -999, -859, 744, -261, 692, -999, -341, -163, -999, -999, 482, -979, 205, -999, 146, 165, 801, 100, -656, 714, -629, 995, 474, 307, -581, -150, -941, -999, -999, -999, -937, -69, -23, 82, -999, -139, -591, -999, -453, -861, -370, -999, -999, -999, 216, 233, -999, 430, -999, 5, -110, -999, -999, -660, 624, -510, -588, -999, -999, 381, -999, 368, 559, -999, 521, -301, -999, 522, 379, -999, -999, -999, -999, 456, 519, -999, -999, 482, 349, -999, -999, 19, -999, -999, 288, -811, -999, -372, -999, -999, -536, -999, -404, -457, -740, 860, -999, -999, -636, -999, -999, 342, -874, -462, -504, 781, 855, -392, -999, -999, -999, 406, -999, -758, 541, -999, -947, -999, -999, -999, -999, -999, -964, -999, 600, -45, -999, -999, -999, -999, -999, -999, -999, -999, -999, -194, -999, -999, -999, -802, -999, -999, -999, -3, -999, -792, 672, 643, -999, 14, -999, -999, 489, 457, -999, -999, -999, -999, 412, -999, 558, -999, -999, -999, -999, -846, 158, -146, -999, -999, -76, -650, -999, -782, -999, -127, -999, -678, -999, -999, -999, -999, -999, -999, -464, -426, -999, -366, -999, -999, -999, -999, -999, 81, -607, 716, -999, -999, -213, -999, 379, -999, -999, -999, -999, 644, 445, -999, -999, -419, -845, -720, -999, -999, -915, -999, -999, -999, -999, -999, -999, -686, 594, -243, -999, 496, -999, 907, -999, -999, -999, -999, -999, -999, 579, 873, 702, -999, -999, -999, -834, -999, -999, -999, -999, -999, -300, -214, -466, -999, -999, 972, -999, -999, -999, 814, -999, -940, -999, 763, -999, -999, -999, -999, -449, -844, -999, -999, -999, -999, -47}
	root := build(nums)
	result := pathSum(root, -243)
	fmt.Println(result)
}
