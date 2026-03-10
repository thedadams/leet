# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        root = None
        if l2 is None:
            return l1
        if l1 is None:
            return l2
        if l1.val < l2.val:
            root = l1
            l1 = l1.next
        else:
            root = l2
            l2 = l2.next
        next = root
        while True:
            if l1 is None:
                next.next = l2
                break
            if l2 is None:
                next.next = l1
                break
            if l2.val < l1.val:
                next.next = l2
                l2 = l2.next
            else:
                next.next = l1
                l1 = l1.next
            next = next.next
        
        return root