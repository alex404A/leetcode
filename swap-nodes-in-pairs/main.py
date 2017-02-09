# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return None
        lists = []
        current = head
        while current:
            lists.append(self.swap(current, current.next))
            current = current.next
        for i in range(len(lists)-1):
            lists[i].next.next = lists[i+1]
        return lists[0]

    def swap(self, node1, node2):
        if node1 and node2:
            node1.next = node2.next
            node2.next = node1
            return node2
        else:
            return node1
