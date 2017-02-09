class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        dp = [[False for i in range(len(p) + 1)] for j in range(len(s) + 1)]
        dp[0][0] = True
        for i in range(1, len(p)+1):
            if p[i-1] == '*':
                if i >= 2:
                    dp[0][i] = dp[0][i-2]
        for i in range(1, len(s)+1):
            for j in range(1, len(p)+1):
                if p[j-1] == '.':
                    dp[i][j] = dp[i-1][j-1]
                elif p[j-1] == '*':
                    dp[i][j] = dp[i][j-2] or (dp[i-1][j] and (
                        s[i-1] == p[j-2] or p[j-2] == '.'))
                else:
                    dp[i][j] = dp[i-1][j-1] and s[i-1] == p[j-1]
        return dp[len(s)][len(p)]


if __name__ == '__main__':
    assert Solution().isMatch("aa", "a") is False
    assert Solution().isMatch("aa", "aa") is True
    assert Solution().isMatch("aaa", "aa") is False
    assert Solution().isMatch("aa", "a*") is True
    assert Solution().isMatch("aa", ".*") is True
    assert Solution().isMatch("ab", ".*") is True
    assert Solution().isMatch("aab", "c*a*b") is True
