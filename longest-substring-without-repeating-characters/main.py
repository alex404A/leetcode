class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        start = end = ans = 0
        charDict = {}
        for char in s:
            charDict[char] = charDict.get(char, 0) + 1
            startOld = start
            while charDict[char] > 1:
                charDict[s[start]] -= 1
                start += 1
            ans = max(ans, end - startOld)
            end += 1
        ans = max(ans, end - start)
        return ans

if __name__ == '__main__':
    solution = Solution()
    print solution.lengthOfLongestSubstring('abcabc')
