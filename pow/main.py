class Solution(object):
    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """

        if n == 0:
            return 1.0
        elif n < 0:
            return 1 / self.myPow(x, 0 - n)
        elif n % 2 == 0:
            return self.myPow(x*x, n/2)
        else:
            return self.myPow(x, n-1) * x

if __name__ == '__main__':
    solution = Solution()
    print(solution.myPow(5.0, -2))
    print(solution.myPow(5.0, 4))
    print(solution.myPow(5.0, 0))
