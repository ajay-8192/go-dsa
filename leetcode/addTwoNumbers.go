package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		var x, y int

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10, Next: nil}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}