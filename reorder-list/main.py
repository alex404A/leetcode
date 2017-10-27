# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def reorderList(self, head):
        """
        :type head: ListNode
        :rtype: void Do not return anything, modify head in-place instead.
        """
        if head == None:
            return
        mid = self.findMid(head)
        second = mid.next
        mid.next = None
        second = self.reverseList(second)
        self.mergeList(head, second)
    
    def findMid(self, head):
        fast = head
        slow = head
        while fast.next is not None and fast.next.next is not None:
            fast = fast.next.next
            slow = slow.next
        return slow

    def reverseList(self, head):
        if head is None:
            return
        n1 = head
        n2 = head.next
        n1.next = None
        while n2 is not None:
            tmp = n2.next
            n2.next = n1
            n1 = n2
            n2 = tmp
        return n1

    def mergeList(self, first, second):
        while first is not None and second is not None:
            tp1 = first.next
            first.next = second
            tp2 = second.next
            second.next = tp1
            first = tp1
            second = tp2
    
    def printList(self, head):
        l = []
        while head is not None:
            l.append(head.val)
            head = head.next
        print(l)

if __name__ == '__main__':
    solution = Solution()
    nodeList = []
    for i in range(4):
        nodeList.append(ListNode(i))
    for i in range(len(nodeList) - 1):
        nodeList[i].next = nodeList[i + 1]

    solution.printList(nodeList[0])
    solution.reorderList(nodeList[0])
    solution.printList(nodeList[0])

