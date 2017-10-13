# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def trimBST(self, root, L, R):
        """
        :type root: TreeNode
        :type L: int
        :type R: int
        :rtype: TreeNode
        """
        if root is None:
            return None
        if root.val < L:
            return self.trimBST(root.right, L, R)
        elif root.val > R:
            return self.trimBST(root.left, L, R)
        else:
            root.left = self.trimBST(root.left, L, R)
            root.right = self.trimBST(root.right, L, R)
            return root
    
    def bfs(self, root):
        nodes = []
        nodes.append((root, None))
        while len(nodes) > 0:
            cmb = nodes.pop()
            node = cmb[0]
            parent = cmb[1]
            print(node.val, None if parent is None else parent.val)
            if node.left is not None:
                nodes.append((node.left, node))
            if node.right is not None:
                nodes.append((node.right, node))

if __name__ == '__main__':
    solution = Solution()
    root = TreeNode(1)
    root.left = TreeNode(0)
    root.right = TreeNode(2)
    L = 1
    R = 2
    trimmer = solution.trimBST(root, L, R)
    solution.bfs(trimmer)