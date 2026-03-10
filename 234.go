/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    _, ok := isPalindromeToEnd(head, head.Next)
    return ok
}

func isPalindromeToEnd(head, next *ListNode) (*ListNode, bool) {
    if next.Next == nil {
        return head.Next, head.Val == next.Val
    }

    nextHead, ok := isPalindromeToEnd(head, next.Next)
    return nextHead.Next, ok && nextHead.Val == next.Val
}