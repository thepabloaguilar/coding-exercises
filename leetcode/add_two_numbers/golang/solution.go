package golang

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	summation := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			summation += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			summation += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: summation % 10}
		current = current.Next

		summation /= 10
	}

	if summation > 0 {
		current.Next = &ListNode{Val: summation}
	}

	return result.Next
}
