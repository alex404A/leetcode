# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        if len(lists) == 0:
            return None
        else:
            return self.iterGenResult(lists)

    def iterGenResult(self, lists):
        tmp = []
        while len(lists) != 1:
            for i in range(0, len(lists), 2):
                if i < len(lists)-1:
                    tmp.append(self.mergeTwoLists(lists[i], lists[i+1]))
                else:
                    tmp.append(self.mergeTwoLists(lists[i], None))
            lists = tmp
            tmp = []
        return lists[0]

    def mergeTwoLists(self, l1, l2):
        if l1 is None:
            return l2
        if l2 is None:
            return l1
        dummy = ListNode(0)
        tmp = dummy
        while l1 and l2:
            if l1.val <= l2.val:
                tmp.next = l1
                l1 = l1.next
                tmp = tmp.next
            else:
                tmp.next = l2
                l2 = l2.next
                tmp = tmp.next
        if l2 is None:
            tmp.next = l1
        else:
            tmp.next = l2
        return dummy.next
