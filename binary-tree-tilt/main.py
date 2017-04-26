class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def __init__(self):
        self.tilt = 0

    def findTilt(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if root is None:
            return 0
        self.iterTree(root)
        return self.tilt

    def iterTree(self, root):
        leftSum = 0
        rightSum = 0
        leftVal = 0
        rightVal = 0
        if root.left is not None:
            self.iterTree(root.left)
            leftSum = root.left.childSum
            leftVal = root.left.val
        if root.right is not None:
            self.iterTree(root.right)
            rightSum = root.right.childSum
            rightVal = root.right.val
        root.childSum = leftSum + leftVal + rightSum + rightVal
        self.tilt += abs(leftSum + leftVal - rightSum - rightVal)

if __name__ == '__main__':
    left = TreeNode(1)
    right = TreeNode(2)
    root = TreeNode(0)
    root.left = left
    root.right = right
    solution = Solution()
    print(solution.findTilt(root))
