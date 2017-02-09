# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        nodeList = self.getNodeList(head)
        if len(nodeList) < n:
            return head
        elif len(nodeList) == n:
            return head.next
        elif n == 1:
            lastNode = nodeList[len(nodeList)-2]
            lastNode.next = None
            return head
        else:
            nodeBeforeDel = nodeList[len(nodeList)-n-1]
            nodeDel = nodeList[len(nodeList)-n]
            nodeBeforeDel.next = nodeDel.next
            return head

    def getNodeList(self, head):
        nodeList, current = [], head
        while True:
            nodeList.append(current)
            if current.next is None:
                break
            current = current.next
        return nodeList
