package golang

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	for list1 != nil || list2 != nil {
		num1 := 101
		if list1 != nil {
			num1 = list1.Val
		}

		num2 := 101
		if list2 != nil {
			num2 = list2.Val
		}

		if num1 < num2 {
			current.Next = &ListNode{Val: num1}
			list1 = list1.Next
		} else {
			current.Next = &ListNode{Val: num2}
			list2 = list2.Next
		}

		current = current.Next
	}

	return head.Next
}
