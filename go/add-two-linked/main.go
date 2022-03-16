package main

// ListNode from Leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	current := &ListNode{
		Val: 0,
	}
	result := current

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next

		if sum > 9 {
			sum = 1
		} else {
			sum = 0
		}
	}

	if sum != 0 {
		current.Next = &ListNode{
			Val: sum,
		}
	}

	return result.Next
}
