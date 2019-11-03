# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def addOneRow(self, root, v, d):
        """
        :type root: TreeNode
        :type v: int
        :type d: int
        :rtype: TreeNode
        """
        if d == 1:
            node = TreeNode(v)
            node.left = root
            return node
        cnt = 1
        queue = [root]
        while len(queue) > 0 and cnt != d - 1:
            newQueue = list(queue)
            queue = []
            for node in newQueue:
                if node.left is not None:
                    queue.append(node.left)
                if node.right is not None:
                    queue.append(node.right)
            cnt += 1
        newNodes = self.duplicateNodes(len(queue) * 2, v)
        for node in queue:
            left = node.left
            right = node.right
            node.left = newNodes.pop()
            node.right = newNodes.pop()
            node.left.left = left
            node.right.right = right
        return root
    
    def duplicateNodes(self, num, v):
        nodes = []
        for i in range(num):
            nodes.append(TreeNode(v))
        return nodes
