class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """

        init = '1'
        for i in range(n - 1):
            init = self.getNext(init)
        return init

    def getNext(self, s):
        num = s[0]
        count = 0
        result = ''
        for i in range(len(s)):
            if num != s[i]:
                result += str(count) + str(num)
                num = s[i]
                count = 1
            else:
                count += 1
        return result + str(count) + str(num)

if __name__ == '__main__':
    solution = Solution()
    print(solution.countAndSay(1))
    print(solution.countAndSay(2))
    print(solution.countAndSay(3))
    print(solution.countAndSay(4))
    print(solution.countAndSay(5))
