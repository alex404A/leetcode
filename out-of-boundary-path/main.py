class Solution(object):
    def __init__(self):
        self.results = {}

    def findPaths(self, m, n, N, i, j):
        """
        :type m: int
        :type n: int
        :type N: int
        :type i: int
        :type j: int
        :rtype: int
        """
        total = self.find(m, n, N, i, j)
        return total % (pow(10, 9) + 7)

    def find(self, m, n, N, i, j):
        if i == -1 or i == m or j == -1 or j == n:
            return 1
        if N == 0:
            return 0

        total = self.results.get((m, n, N, i, j))
        if total is not None:
            return total

        leftCnt = self.findPaths(m, n, N - 1, i - 1, j)
        rightCnt = self.findPaths(m, n, N - 1, i + 1, j)
        upCnt = self.findPaths(m, n, N - 1, i, j - 1)
        downCnt = self.findPaths(m, n, N - 1, i, j + 1)
        total = leftCnt + rightCnt + upCnt + downCnt
        self.results[(m, n, N, i, j)] = total
        return total

if __name__ == '__main__':
    solution = Solution()
    print(solution.findPaths(8, 50, 23, 5, 26))
