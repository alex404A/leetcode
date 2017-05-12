class Solution(object):
    def __init__(self):
        self.results = {}
        self.divisor = pow(10, 9) + 7

    def findPaths(self, m, n, N, i, j):
        """
        :type m: int
        :type n: int
        :type N: int
        :type i: int
        :type j: int
        :rtype: int
        """
        if i == -1 or i == m or j == -1 or j == n:
            return 1
        if i >= N and m - i - 1 >= N and j >= N and n - j - 1 >= N:
            return 0

        total = self.results.get((N, i, j))
        if total is not None:
            return total

        leftCnt = self.findPaths(m, n, N - 1, i - 1, j)
        rightCnt = self.findPaths(m, n, N - 1, i + 1, j)
        upCnt = self.findPaths(m, n, N - 1, i, j - 1)
        downCnt = self.findPaths(m, n, N - 1, i, j + 1)
        total = (leftCnt + rightCnt + upCnt + downCnt) % self.divisor
        self.results[(N, i, j)] = total
        return total

if __name__ == '__main__':
    solution = Solution()
    print(solution.findPaths(8, 7, 16, 1, 5))
