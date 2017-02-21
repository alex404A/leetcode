class Solution(object):
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        needleLen = len(needle)
        haystackLen = len(haystack)
        if needleLen == 0:
           return 0 
        for i in range(0, haystackLen):
            k, j, isNeedleDetected = i, 0, False
            if haystackLen - i >= needleLen:
                while j < needleLen:
                    if needle[j] == haystack[k]:
                        if j == needleLen - 1:
                            isNeedleDetected = True
                        j += 1
                        k += 1
                    else:
                        break
                if isNeedleDetected is True:
                    return i
            else:
                return -1
        return -1

if __name__ == '__main__':
    solution = Solution()
    print(solution.strStr('aabbaa', 'bb'))
    print(solution.strStr('aababaa', 'bb'))
    print(solution.strStr('', 'bb'))
