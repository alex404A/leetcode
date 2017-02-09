class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        if n == 0:
            return []
        return self.recurseGenList(n, n)

    def recurseGenList(self, left, right):
        if left == 1 and right == 1:
            return ['()']
        elif left == 0 and right == 1:
            return [')']
        else:
            tmp1 = tmp2 = []
            if left >= 1:
                tmp1 = self.genList('(', self.recurseGenList(left-1, right))
            if left < right:
                tmp2 = self.genList(')', self.recurseGenList(left, right-1))
            return tmp1 + tmp2

    def genList(self, left, parenthesesList):
        return [left + p for p in parenthesesList]


if __name__ == '__main__':
    solution = Solution()
    print(solution.generateParenthesis(3))
