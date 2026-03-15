package golang

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if seen[node] {
			return true
		}

		seen[node] = true
		node = node.Next
	}

	return false
}
