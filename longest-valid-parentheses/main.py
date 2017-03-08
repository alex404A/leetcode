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
                leftStack.append(i)
                continuousRightCnt = 0
            else if s[i] == ')':
                if (continuousRightCnt == len(self.leftStack)):
                    self.popLastNElementsAndUpdate(continuousRightCnt)
                    continuousRightCnt = 0
                else:
                    rightStack.append(i)
                    continuousRightCnt += 1

        return self.maxLen

    def popLastNElementsAndUpdate(self, n):
        if n = 0:
            return
        leftIndex, rightIndex, newInterval
        for i in range(0, n):
            leftIndex = self.leftStack.pop()
            rightIndex = self.rightStack.pop()
        leftInterval = rightIndexToIntervalDict.pop(leftIndex - 1, index)
        if leftInterval is not None:
            newInterval = (leftInterval[0], rightIndex)
        else:
            newInterval = (leftIndex, rightIndex)
        rightIndexToIntervalDict[rightIndex] = newInterval
        if newInterval[1] - newInterval[0] + 1 > self.maxLen:
            self.maxLen = newInterval[1] - newInterval[0] + 1
