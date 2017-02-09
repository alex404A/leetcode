class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        result = ''
        for index in range(len(s)):
            longestEven = self.getLongestPalindDrome(s, index, index+1)
            if len(longestEven) > len(result):
                result = longestEven
            longestOdd = self.getLongestPalindDrome(s, index, index)
            if len(longestOdd) > len(result):
                result = longestOdd
        return result

    def getLongestPalindDrome(self, s, left, right):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return s[left+1:right]

if __name__ == '__main__':
    solution = Solution()
    print solution.longestPalindrome('53466427')
