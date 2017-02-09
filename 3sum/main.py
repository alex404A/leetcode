class Solution(object):
    def binarySearch(self, array, item):
        if len(array) == 0:
            return False
        if len(array) == 1:
            if array[0] != item:
                return False
            else:
                return True
        left = 0
        right = len(array) - 1
        middle = (left + right) / 2
        if array[middle] == item:
            return True
        elif array[middle] > item:
            return self.binarySearch(array[:middle], item)
        else:
            return self.binarySearch(array[middle+1:], item)

    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        negList = [x for x in nums if x < 0]
        posList = [x for x in nums if x >= 0]
        results = []

        lastNegNum = 0
        for negNumIndex, negNum in enumerate(negList):
            if negNum == lastNegNum:
                continue
            else:
                lastNegNum = negNum
                lastPosNum = -1
                for posNumIndex, posNum in enumerate(posList):
                    if lastPosNum == posNum:
                        continue
                    else:
                        lastPosNum = posNum
                        remain = 0 - (negNum + posNum)
                        if remain < 0:
                            if self.binarySearch(
                                negList[negNumIndex + 1:],
                                remain
                            ):
                                results.append([negNum, remain, posNum])
                        else:
                            if self.binarySearch(
                                posList[posNumIndex + 1:],
                                remain
                            ):
                                results.append([negNum, posNum, remain])

        zeroList = [x for x in posList if x == 0]
        if len(zeroList) >= 3:
            results.append([0, 0, 0])
        return results

if __name__ == '__main__':
    solution = Solution()
    test1 = [-1, 0, 1, 0]
    print(solution.threeSum(test1))
