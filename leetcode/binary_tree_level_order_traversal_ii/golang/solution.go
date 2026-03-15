package golang

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	values := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		levelValues := make([]int, 0, size)
		for idx := range size {
			node := queue[idx]
			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		values = append(values, levelValues)
		queue = queue[size:]
	}

	slices.Reverse(values)
	return values
}
