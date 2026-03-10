/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	third := head.Next.Next
    head.Next = nil

	for third != nil {
        second.Next = first

        first = second
        second = third
        third = third.Next
	}

    second.Next = first

	return second
}