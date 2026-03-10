/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    item := head
    var prev *ListNode
    for item != nil {
        if item.Val == val {
            if prev != nil {
                prev.Next = item.Next
            } else {
                head = item.Next
            }
        } else {
            prev = item
        }
        
        item = item.Next
    }

    return head
}