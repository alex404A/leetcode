class Solution(object):
    def __init__(self):
        self.leftStack = []
        self.rightStack = []
        self.rightIndexToIntervalDict = {}
        self.maxLen = 0

    def longestValidParentheses(self, s):
        """
        :type s: str
        :rtype: int
        """

        continuousRightCnt = 0

        for i in range(0, len(s)):
            if s[i] == '(':
                self.popLastNElementsAndUpdate(continuousRightCnt)
                self.leftStack.append(i)
                continuousRightCnt = 0
            elif s[i] == ')':
                if (continuousRightCnt == len(self.leftStack)):
                    self.popLastNElementsAndUpdate(continuousRightCnt)
                    continuousRightCnt = 0
                else:
                    self.rightStack.append(i)
                    continuousRightCnt += 1

        self.popLastNElementsAndUpdate(continuousRightCnt)
        return self.maxLen

    def popLastNElementsAndUpdate(self, n):
        if n == 0:
            return
        rightIndex = self.rightStack[len(self.rightStack) - 1]
        for i in range(0, n):
            leftIndex = self.leftStack.pop()
            self.rightStack.pop()
        leftInterval = self.rightIndexToIntervalDict.pop(leftIndex - 1, None)
        if leftInterval is not None:
            newInterval = (leftInterval[0], rightIndex)
        else:
            newInterval = (leftIndex, rightIndex)
        self.rightIndexToIntervalDict[rightIndex] = newInterval
        if newInterval[1] - newInterval[0] + 1 > self.maxLen:
            self.maxLen = newInterval[1] - newInterval[0] + 1

if __name__ == '__main__':
    solution = Solution()
    s = "))(()()))"
    solution.longestValidParentheses(s)
