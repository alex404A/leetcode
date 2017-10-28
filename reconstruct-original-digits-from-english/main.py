class Solution(object):
    def originalDigits(self, s):
        """
        :type s: str
        :rtype: str
        """
        statistics = self.collectStatistics(s)
        return self.collectInfo(statistics)
    
    def collectInfo(self, statistics):
        infoList = [0] * 10
        cnt = statistics[self.getPos('z')]
        infoList[0] = self.deductStatistics('z', ['z', 'e', 'r', 'o'], statistics)
        infoList[2] = self.deductStatistics('w', ['t', 'w', 'o'], statistics)
        infoList[4] = self.deductStatistics('u', ['f', 'o', 'u', 'r'], statistics)
        infoList[6] = self.deductStatistics('x', ['s', 'i', 'x'], statistics)
        infoList[8] = self.deductStatistics('g', ['e', 'i', 'g', 'h', 't'], statistics)
        infoList[7] = self.deductStatistics('s', ['s', 'e', 'v', 'e', 'n'], statistics)
        infoList[5] = self.deductStatistics('v', ['f', 'i', 'v', 'e'], statistics)
        infoList[1] = self.deductStatistics('o', ['o', 'n', 'e'], statistics)
        infoList[3] = self.deductStatistics('t', ['t', 'h', 'r', 'e', 'e'], statistics)
        infoList[9] = self.deductStatistics('i', ['n', 'i', 'n', 'e'], statistics)
        result = ''
        for index, val in enumerate(infoList):
            result += str(index) * val
        return result

    
    def collectStatistics(self, s):
        statistics = [0] * 26
        for i in range(len(s)):
            statistics[self.getPos(s[i])] += 1
        return statistics

    def deductStatistics(self, label, charList, statistics):
        cnt = statistics[self.getPos(label)]
        if cnt != 0:
            for char in charList:
                statistics[self.getPos(char)] -= cnt
            return cnt
        else:
            return 0

    def getPos(self, char):
        return ord(char) - ord('a')

if __name__ == '__main__':
    solution = Solution()
    print(solution.originalDigits('nnei'))

        