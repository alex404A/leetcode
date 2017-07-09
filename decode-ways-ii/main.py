class Solution(object):
    def numDecodings(self, s):
        """
        :type s: str
        :rtype: int
        """
        sLen = len(s)
        if sLen == 0:
            return 0
        elif sLen == 1:
            return self.getOneNumCmb(s[0])
        elif sLen == 2:
            return self.getTwoNumsCmb(s[0], s[1]) + self.getOneNumCmb(s[0]) * self.getOneNumCmb(s[1])

        nextToLastCnt = self.getOneNumCmb(s[0])
        lastCnt = self.getTwoNumsCmb(s[0], s[1]) + nextToLastCnt * self.getOneNumCmb(s[1])
        divisor = pow(10, 9) + 7

        for i in range(2, sLen):
            tmpCnt = (self.getOneNumCmb(s[i]) * lastCnt + self.getTwoNumsCmb(s[i - 1], s[i]) * nextToLastCnt) % divisor
            nextToLastCnt = lastCnt
            lastCnt = tmpCnt

        return lastCnt

    def getOneNumCmb(self, numStr):
        return 9 if numStr == '*' else 1 if numStr != '0' else 0

    def getTwoNumsCmb(self, numStr1, numStr2):
        if numStr1 == '*':
            if numStr2 == '*':
                return 15
            else:
                return 2 if int(numStr2) <= 6 else 1
        else:
            num1 = int(numStr1)
            if num1 == 0 or num1 > 2:
                return 0
            elif num1 == 1:
                return 9 if numStr2 == '*' else 1
            else:
                if numStr2 == '*':
                    return 6
                else:
                    return 1 if int(numStr2) <= 6 else 0


if __name__ == '__main__':
    solution = Solution()
    print(solution.numDecodings('4*28*'))
