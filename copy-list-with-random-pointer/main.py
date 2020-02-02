class Node:
    def __init__(self, x, next=None, random=None):
        self.val = int(x)
        self.next = next
        self.random = random

class Solution(object):

    def __init__(self):
        self.dict = {}
        self.head = None

    def copyRandomList(self, head):
        if head == None:
            return None
        self.copyAndStore(head)
        self.replaceRndom()
        return self.head
        """
        :type head: Node
        :rtype: Node
        """
    
    def copyAndStore(self, head):
        """
        :type head: Node
        :rtype: list
        """
        copiedNode = self.copy(head, None)
        self.dict[head.val] = []
        self.dict[head.val].append((copiedNode, head))
        self.head = copiedNode
        head = head.next
        previous = copiedNode
        while head != None:
            copiedNode = self.copy(head, previous)
            if head.val not in self.dict:
                self.dict[head.val] = []
            self.dict[head.val].append((copiedNode, head))
            previous = copiedNode
            head = head.next
    
    def replaceRndom(self):
        for val in self.dict:
            for t in self.dict[val]:
                random = t[0].random
                if random is not None:
                    for candidate in self.dict[random.val]:
                        if candidate[1] == random:
                            t[0].random = candidate[0]
                            break

    def copy(self, node, previous):
        """
        :type node: Node
        :type previous: Node
        :rtype: Node
        """
        node = Node(node.val, None, node.random)
        if previous != None:
            previous.next = node
        return node

def test(head):
    list = []
    while head is not None:
        random = -1 if head.random is None else head.random.val
        list.append((head.val, random))
        head = head.next
    print(list)

if __name__ == '__main__':
    solution = Solution()
    head = Node(0, None, None)
    head.next = Node(1, None, head)
    copiedNode = solution.copyRandomList(head)
    test(copiedNode)
