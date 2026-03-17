package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, level int) int {
	if node == nil {
		return level
	}

	return max(
		dfs(node.Left, level+1),
		dfs(node.Right, level+1),
	)
}
