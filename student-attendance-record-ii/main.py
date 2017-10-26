"""
d[i] = the num of available solutions in len i without A
without A:
1. ends with P
    d[i - 1]
2. ends with L
    PL d[i - 2]
    PLL d[i - 3]
    d[i - 2] + d[i - 3]
d[i] = d[i - 1] + d[i - 2] + d[i - 3]
with A:
"""
class Solution(object):
    def checkRecord(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n == 0:
            return 0
        elif n == 1:
            return 3
        elif n == 2:
            return 8
        elif n == 3:
            return 19
        result = 0
        M = 1000000007
        d = [0] * n
        d[0] = 1
        d[1] = 2
        d[2] = 4
        d[3] = 7
        i = 4
        while i < n:
            d[i] = (d[i - 1] + d[i - 2] + d[i - 3]) % M
            i += 1
        result = (d[n - 1] + d[n - 2] + d[n - 3]) % M
        print(result)
        for i in range(n):
            result += (d[i] * d[n - 1 - i])
            result %= M
        return result

if __name__ == '__main__':
    solution = Solution()
    print(solution.checkRecord(4))
        