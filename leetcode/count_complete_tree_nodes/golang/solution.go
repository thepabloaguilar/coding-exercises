package golang

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := getLeftDepth(root.Left)
	rightDepth := getRightDepth(root.Right)

	if leftDepth == rightDepth {
		return int(math.Pow(2, float64(leftDepth))) - 1
	}

	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}

func getLeftDepth(root *TreeNode) int {
	depth := 1
	for root != nil {
		depth++
		root = root.Left
	}
	return depth
}

func getRightDepth(root *TreeNode) int {
	depth := 1
	for root != nil {
		depth++
		root = root.Right
	}
	return depth
}
