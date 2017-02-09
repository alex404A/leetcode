# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if not head:
            return None
        if k == 1:
            return head
        nodeList = []
        result = []
        current = head
        while current:
            nodeList.append(current)
            current = current.next
        for i in range(0, len(nodeList), k):
            result.append(self.reverseGroup(nodeList[i], k))
        print(result)
        for i in range(len(result)-1):
            result[i][1].next = result[i+1][0]
        return result[0][0]

    def reverseGroup(self, node, k):
        head = node
        nodeList = []
        flag = True
        for i in range(k):
            nodeList.append(node)
            if node.next:
                node = node.next
            else:
                if i < k-1:
                    flag = False
                    break
        if not flag:
            return [head, node]
        for i in range(k-1, 0, -1):
            nodeList[i].next = nodeList[i-1]
        nodeList[0].next = None
        return [nodeList[k-1], nodeList[0]]

if __name__ == '__main__':
    solution = Solution()
    node1 = ListNode(1)
    node2 = ListNode(2)
    node3 = ListNode(3)
    node1.next = node2
    node2.next = node3
    current = solution.reverseKGroup(node1, 2)
    while current:
        print(str(current.val) + '->'),
        current = current.next
