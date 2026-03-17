package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return traverse(root.Left, root.Right)
}

func traverse(nodeA *TreeNode, nodeB *TreeNode) bool {
	if nodeA == nil && nodeB == nil {
		return true
	}
	if nodeA == nil || nodeB == nil {
		return false
	}
	if nodeA.Val != nodeB.Val {
		return false
	}

	return traverse(nodeA.Left, nodeB.Right) && traverse(nodeA.Right, nodeB.Left)
}
