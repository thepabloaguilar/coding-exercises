package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfLevels(root *TreeNode) []float64 {
	averageByLevel := make([]float64, 0)

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0
		size := len(queue)
		for idx := range size {
			node := queue[idx]
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		averageByLevel = append(averageByLevel, float64(sum)/float64(size))
		queue = queue[size:]
	}

	return averageByLevel
}
