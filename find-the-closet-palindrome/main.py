class Solution(object):
    def nearestPalindromic(self, num):
        """
        :type num: str
        :rtype: str
        """
        def getBoundary(realNum, interval, numLen):
            palLen = len(str(pal))
            if palLen == numLen:
                return pal
            elif palLen > numLen:
                result = 0
                result += 1
                if numLen >= 1:
                    result += pow(10, numLen) * 1
                    return result
            else:
                result = 0
                for i in range(palLen):
                    result += pow(10, i) * 9
                return result

        def getIntervals(numLen):
            pass

        numLen = len(num)
        if numLen >= 2 and num[0] == '1' and num[1:] == '0' * (numLen - 1):
            return '9' * (numLen - 1)

        interval = 0
        quotient = numLen / 2
        if numLen % 2 == 0:
            interval = pow(10, quotient - 1) + pow(10, quotient)
        else:
            interval = pow(10, quotient)

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
            firstPal = getBoundary(realNum + interval, numLen)
            secondPal = getBoundary(realNum - interval, numLen)
        elif realNum > firstPal:
            secondPal = getBoundary(firstPal + interval, numLen)
        else:
            secondPal = getBoundary(firstPal - interval, numLen)
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
    print(solution.nearestPalindromic('11011'))
