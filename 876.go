/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	_, m := findMiddle(head, 1)
	return m
}

func findMiddle(head *ListNode, idx int) (int, *ListNode) {
	if head == nil {
		return idx + 1, nil
	}

	l, m := findMiddle(head.Next, idx+1)
	if m != nil {
		return l, m
	}

	if l/2 == idx {
		return l, head
	} else if (l-1)/2 == idx {
		return l, head
	}

	return l, nil
}
