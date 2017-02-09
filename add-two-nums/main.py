# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        l1Next = l1
        l2Next = l2
        result = ListNode(0)
        resultNext = result
        while True:
            if (l1Next is None) and (l2Next is None):
                tmpNext = self.genNode(0, 0, resultNext, True)
                break
            elif (l1Next is None) and (l2Next is not None):
                tmpNext = self.genNode(0, l2Next.val, resultNext)
                resultNext = tmpNext
                l2Next = l2Next.next
                l1Next = None
            elif (l1Next is not None) and (l2Next is None):
                tmpNext = self.genNode(l1Next.val, 0, resultNext)
                resultNext = tmpNext
                l1Next = l1Next.next
                l2Next = None
            else:
                tmpNext = self.genNode(l1Next.val, l2Next.val, resultNext)
                resultNext = tmpNext
                l1Next = l1Next.next
                l2Next = l2Next.next
        self.delLastNode(result)
        return result

    def genNode(self, val1, val2, node, flag=False):
        val = (val1 + val2 + node.val) % 10
        carry = (val1 + val2 + node.val) / 10
        node.val = val
        if not flag:
            nextNode = ListNode(carry)
            node.next = nextNode
            return nextNode
        else:
            if val == 1:
                nextNode = ListNode(carry)
                node.next = nextNode
                return nextNode
            else:
                node.next = None
                return None

    def delLastNode(self, firstNode):
        nextNode = firstNode.next
        lastNode = firstNode
        if nextNode is None:
            return 0
        while True:
            print lastNode.val, nextNode.val
            if nextNode.next is None and nextNode.val == 0:
                lastNode.next = None
                break
            else:
                lastNode = nextNode
                nextNode = nextNode.next


def genNodesList(numList):
    node = ListNode(0)
    nextNode = node
    i = 0
    while True:
        nextNode.val = numList[i]
        i += 1
        if i == len(numList):
            nextNode.next = None
            break
        else:
            nextNode.next = ListNode(0)
            nextNode = nextNode.next
    return node

if __name__ == "__main__":
    a1 = genNodesList([0])
    b1 = genNodesList([0])
    solution = Solution()
    result = solution.addTwoNumbers(a1, b1)
    resultNext = result
    text = ''
    while True:
        text += str(resultNext.val)
        if resultNext.next is not None:
            text += '-->'
            resultNext = resultNext.next
        else:
            break
    print text
