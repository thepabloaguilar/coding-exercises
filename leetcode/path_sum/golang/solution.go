package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}
