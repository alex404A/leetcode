class Solution(object):
    def thirdMax(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        def resetMaximumNum(num):
            if num > maximumNums[0]:
                maximumNums[2] = maximumNums[1]
                maximumNums[1] = maximumNums[0]
                maximumNums[0] = num
            elif num < maximumNums[0]:
                if num > maximumNums[1]:
                    maximumNums[2] = maximumNums[1]
                    maximumNums[1] = num
                elif num < maximumNums[1]:
                    if num > maximumNums[2]:
                        maximumNums[2] = num

        negInfinity = float('-infinity')
        maximumNums = [negInfinity, negInfinity, negInfinity]
        for index, value in enumerate(nums):
            resetMaximumNum(value)
        if negInfinity == maximumNums[0]:
            return negInfinity
        else:
            return maximumNums[2] if negInfinity != maximumNums[2]\
                else maximumNums[0]


if __name__ == "__main__":
    testList = [1, 2, 2, 5, 3, 5]
    solution = Solution()
    print solution.thirdMax(testList)
