package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetMinimumDifference(root *TreeNode) int {
	const maxDiff = 1<<31 - 1 // Use a large but safe number
	return traverse(root, -maxDiff, maxDiff)
}

func traverse(node *TreeNode, low int, high int) int {
	if node == nil {
		return high - low
	}

	left := traverse(node.Left, low, node.Val)
	right := traverse(node.Right, node.Val, high)

	if left < right {
		return left
	}
	return right
}
