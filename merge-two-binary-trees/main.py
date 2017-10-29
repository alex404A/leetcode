# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def mergeTrees(self, t0, t1):
        """
        :type t1: TreeNode
        :type t2: TreeNode
        :rtype: TreeNode
        """
        if t0 is None:
            return t1
        if t1 is None:
            return t0
        queue0 = []
        queue1 = []
        root = TreeNode(0)
        queue0.append((t0, t1))
        queue1.append(root)
        while len(queue1) > 0:
            cmb = queue0.pop(0)
            node = queue1.pop(0)
            node.val = cmb[0].val + cmb[1].val
            left0 = cmb[0].left
            left1 = cmb[1].left
            right0 = cmb[0].right
            right1 = cmb[1].right
            if left0 is None or left1 is None:
                node.left = left0 if left1 is None else left1
            else:
                queue0.append((left0, left1))
                leftNode = TreeNode(0)
                node.left = leftNode
                queue1.append(leftNode)
            if right0 is None or right1 is None:
                node.right = right0 if right1 is None else right1
            else:
                queue0.append((right0, right1))
                rightNode = TreeNode(0)
                node.right = rightNode
                queue1.append(rightNode)
        return root

    def iterTree(self, root):
        queue = []
        if root is None:
            return
        queue.append(root)
        while len(queue) > 0:
            node = queue.pop(0)
            print(node.val)
            if node.left is not None:
                queue.append(node.left)
            if node.right is not None:
                queue.append(node.right)

if __name__ == '__main__':
    a0 = TreeNode(1)
    a1 = TreeNode(3)
    a2 = TreeNode(2)
    a3 = TreeNode(5)
    a0.left = a1
    a0.right = a2
    a1.left = a3
    b0 = TreeNode(2)
    b1 = TreeNode(1)
    b2 = TreeNode(3)
    b3 = TreeNode(4)
    b4 = TreeNode(7)
    b0.left = b1
    b0.right = b2
    b1.right = b3
    b2.right = b4
    solution = Solution()
    root = solution.mergeTrees(a0, b0)
    solution.iterTree(root)
        