class Solution(object):
    def nearestPalindromic(self, num):
        """
        :type num: str
        :rtype: str
        """
        def getBoundary(realNum, intervals, isIncreased):
            boundaries = []
            for interval in intervals:
                newBoundary = realNum + interval if isIncreased else realNum - interval
                if isPalindrome(newBoundary):
                    boundaries.append(newBoundary)
            result = -1
            maxInterval = 2147483647
            for boundary in boundaries:
                interval = abs(realNum - boundary)
                if interval < maxInterval:
                    result = boundary
                    maxInterval = interval
            return result

        def getIntervalsForOdd(numLen):
            results = []
            first = '1' * numLen
            second = '1' * (numLen / 2) + '2' + '1' * (numLen / 2)
            for i in range(numLen / 2 + 1):
                results.append(int(second) - int(first))
                lenForMidInFirst = 2 * i + 1
                lenForSideInFirst = (numLen - lenForMidInFirst) / 2
                first = '1' * lenForSideInFirst + '9' * lenForMidInFirst + '1' * lenForSideInFirst
                lenForMidInSecond = lenForMidInFirst
                lenForOneInSecond = (numLen - lenForMidInSecond - 2) / 2
                second = '1' * lenForOneInSecond + '2' + '0' * lenForMidInSecond + '2' + '1' * lenForOneInSecond
            results.append(2)
            results.append(1)
            return results

        def getIntervalsForEven(numLen):
            results = []
            first = '1' * numLen
            second = '1' * (numLen / 2 - 1) + '22' + '1' * (numLen / 2 - 1)
            for i in range(numLen / 2):
                results.append(int(second) - int(first))
                lenForMidInFirst = 2 * (i + 1)
                lenForSideInFirst = (numLen - lenForMidInFirst) / 2
                first = '1' * lenForSideInFirst + '9' * lenForMidInFirst + '1' * lenForSideInFirst
                lenForMidInSecond = lenForMidInFirst
                lenForOneInSecond = (numLen - lenForMidInSecond - 2) / 2
                second = '1' * lenForOneInSecond + '2' + '0' * lenForMidInSecond + '2' + '1' * lenForOneInSecond
            results.append(2)
            results.append(1)
            return results

        def isPalindrome(num):
            numStr = str(num)
            for i in range(len(numStr)):
                j = len(numStr) - i - 1
                if j <= i:
                    return True
                if numStr[i] != numStr[j]:
                    return False

        numLen = len(num)
        intervals = getIntervalsForOdd(numLen) if numLen % 2 else getIntervalsForEven(numLen)
        print(intervals)

        firstPal = 0
        for i in range(numLen):
            j = numLen - 1 - i
            if j < i:
                break
            elif j > i:
                firstPal += pow(10, i) * int(num[i])
                firstPal += pow(10, j) * int(num[i])
            else:
                firstPal += pow(10, i) * int(num[i])

        realNum = int(num)
        secondPal = 0
        print(firstPal)
        if realNum == firstPal:
            firstPal = getBoundary(realNum, intervals, True)
            secondPal = getBoundary(realNum, intervals, False)
        elif realNum > firstPal:
            secondPal = getBoundary(firstPal, intervals, True)
        else:
            secondPal = getBoundary(firstPal, intervals, False)
        print(firstPal, secondPal)

        diff1 = abs(realNum - firstPal)
        diff2 = abs(realNum - secondPal)
        if diff1 < diff2:
            return str(firstPal)
        elif diff2 < diff1:
            return str(secondPal)
        else:
            return str(min(firstPal, secondPal))

if __name__ == '__main__':
    solution = Solution()
    print(solution.nearestPalindromic('2'))
