/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    var (
        carry int
        prevNext1, prevNext2 *ListNode
        next1, next2 = l1, l2
    )
    for {
        if next1 == nil {
            if next2 == nil {
                if carry > 0 {
                    prevNext1.Next = &ListNode{Val: carry}
                }
                return l1
            }
            next1 = next2
            next2 = nil
            prevNext1 = prevNext2
            prevNext2 = nil
            l1 = l2
        }
        if next2 != nil {
            next1.Val += next2.Val
        }

        next1.Val += carry
        carry = next1.Val / 10
        next1.Val %= 10
        
        if next2 != nil {
            next2.Val = next1.Val
            prevNext2 = next2
            next2 = next2.Next
        }
        
        prevNext1 = next1
        next1 = next1.Next
    }
}