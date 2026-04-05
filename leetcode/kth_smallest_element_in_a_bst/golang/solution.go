package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	result := 0
	inorder(root, k, new(0), &result)
	return result
}

func inorder(node *TreeNode, k int, count *int, result *int) bool {
	if node == nil {
		return false
	}

	if inorder(node.Left, k, count, result) {
		return true
	}

	*count++
	if *count == k {
		*result = node.Val
		return true
	}

	return inorder(node.Right, k, count, result)
}
