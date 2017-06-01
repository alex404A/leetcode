class Solution(object):
    def findIntegers(self, num):
        """
        :type num: int
        :rtype: int
        """
        binNum = bin(num)[2:]
        count = 0
        count += self.getCountLessThanAndEqualToLen(len(binNum) - 1)
        print("step1: " + str(count))
        isBinNumValid, posList = self.parseBinNum(binNum)
        if isBinNumValid is True:
            count += 1
        print("step2: " + str(count))
        for i in posList:
            count += self.getCountLessThanAndEqualToLen(i - 1)
        return count

    def getCountLessThanAndEqualToLen(self, length):
        countInDiffBits = [
            1, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233,
            377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711,
            28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040
        ]
        count = 0
        for i in range(length + 1):
            count += countInDiffBits[i]
        return count

    def parseBinNum(self, binNum):
        binNumLen = len(binNum)
        posList = []
        isLastNumOne = True
        for i, num in enumerate(binNum[1:]):
            if num == '0':
                isLastNumOne = False
            else:
                posList.append(binNumLen - i - 1)
                if isLastNumOne is True:
                    return False, posList
                else:
                    isLastNumOne = True
        return True, posList

if __name__ == '__main__':
    solution = Solution()
    print('step3: ' + str(solution.findIntegers(11)))
