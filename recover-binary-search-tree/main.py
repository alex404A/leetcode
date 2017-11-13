# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):

    def __init__(self):
        self.sortedNodes = []
        self.preNode = None
        self.isFirstNode = True
        self.wrongNodes = []

    def recoverTree(self, root):
        """
        :type root: TreeNode
        :rtype: void Do not return anything, modify root in-place instead.
        """
        if root is None:
            return
        self.iterate(root, None)
        if len(self.wrongNodes) == 0:
            return
        firstNode = self.wrongNodes[0][0]
        firstParent = self.wrongNodes[0][1]
        if len(self.wrongNodes) == 1:
            self.swap(firstNode, firstParent)
            return
        secondNode = self.wrongNodes[1][0]
        secondParent = self.wrongNodes[1][1]
        self.swap(firstNode, secondNode)
            
    def iterate(self, root, parent):
        if root.left is not None:
            self.iterate(root.left, root)
        if self.preNode is not None and self.preNode.val > root.val:
            if self.isFirstNode:
                self.wrongNodes.append((self.preNode, root))
                self.isFirstNode = False
            else:
                self.wrongNodes.append((root, parent))
        self.preNode = root
        if root.right is not None:
            self.iterate(root.right, root)
    
    def swap(self, first, second):
        tmp = first.val
        first.val = second.val
        second.val = tmp
    
    def iteratePrint(self, root):
        if root.left is not None:
            self.iteratePrint(root.left)
        print(root.val)
        if root.right is not None:
            self.iteratePrint(root.right)

if __name__ == '__main__':
    solution = Solution()
    parent = TreeNode(0)
    left = TreeNode(1)
    print(parent)
    print(left)
    parent.left = left
    solution.recoverTree(parent)
    print(parent)
    solution.iteratePrint(parent)

        