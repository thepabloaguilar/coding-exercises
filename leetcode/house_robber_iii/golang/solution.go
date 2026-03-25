package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Rob(root *TreeNode) int {
	return max(calculateRob(root))
}

func calculateRob(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftNoRob, leftRob := calculateRob(root.Left)
	rightNoRob, rightRob := calculateRob(root.Right)

	return max(leftRob, leftNoRob) + max(rightRob, rightNoRob),
		root.Val + leftNoRob + rightNoRob
}
