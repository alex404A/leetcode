class Solution(object):
    def isPalindrome(self, num):
        """
        :type num: int
        :rtype: bool
        """
        if num < 0:
            return False
        strNum = str(num)
        numLen = len(strNum)
        for i in range(numLen):
            j = numLen - i - 1
            if i >= j:
                return True
            if strNum[i] != strNum[j]:
                return False

if __name__ == '__main__':
    solution = Solution()
    print(solution.isPalindrome(5665))
    print(solution.isPalindrome(-5665))
    print(solution.isPalindrome(0))
    print(solution.isPalindrome(56665))
    print(solution.isPalindrome(1234421))
