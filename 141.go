/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
   if head == nil || head.Next == nil {
    return false
   } 

   fast := head.Next.Next

   for head.Next != nil && fast != nil {
    if head == fast {
        return true
    }

    if fast.Next == nil {
        return false
    }

    head = head.Next
    fast = fast.Next.Next
   }

    return false
}