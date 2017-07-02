class Solution(object):
    def smallestRange(self, numsList):
        """
        :type numsList: List[List[int]]
        :rtype: List[int]
        """

        start = maxStart = max([nums[0] for nums in numsList])
        end = minEnd = min([nums[len(nums) - 1] for nums in numsList])
        intervalList = []
        for nums in numsList:
            startIndex = self.biSearch(maxStart, nums)
            endIndex = self.biSearch(minEnd, nums)
            if endIndex != len(nums) - 1 and endIndex != -1:
                endIndex += 1
            intervalList.append([startIndex, endIndex])
            if startIndex == -1 or endIndex == -1:
                return []
            if nums[startIndex] < maxStart:
                start = nums[startIndex]
            if nums[endIndex] > minEnd:
                end = nums[endIndex]
        comb = {
            "interval": (start, maxStart) if maxStart - start <= end - endIndex else (minEnd, end),
            "length": min(maxStart - start, end - minEnd)
        }
        print(intervalList)

        commonNumDict = {}
        for i, interval in enumerate(intervalList):
            nums = numsList[i]
            startIndex = interval[0]
            endIndex = interval[1]
            for j in range(startIndex, endIndex + 1):
                if commonNumDict.get(nums[j]) is  None:
                    commonNumDict[nums[j]] = set()
                commonNumDict[nums[j]].add(i)

        validNums = [num for num in commonNumDict.iterkeys()]
        print(validNums)

        print(commonNumDict)

    def biSearch(self, target, nums):
        left = 0
        right = len(nums) - 1
        while left < right:
            middle = (left + right) / 2
            if target >= nums[middle] and target < nums[middle + 1]:
                return middle
            elif target < nums[middle]:
                right = middle
            elif target >= nums[middle + 1]:
                left = middle + 1
        return left if target >= nums[left] else -1

if __name__ == '__main__':
    solutioon = Solution()
    numsList = [
        [4,10,15,24,26],
        [0,9,12,20],
        [5,18,22,30]
    ]
    solutioon.smallestRange(numsList)
