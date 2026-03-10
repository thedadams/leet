# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        next = head
        while next is not None and next.next is not None:
            if next.val == next.next.val:
                next.next = next.next.next
            else:
                next = next.next
        return head