class Solution(object):
    def __init__(self):
        self.numToCntDict = {}

    def triangleNumber(self, duplicateNums):
        """
        :type nums: List[int]
        :rtype: int
        """
        self.genNumToCntDict(duplicateNums)
        self.nums = [num for num in self.numToCntDict.iterkeys()]
        self.nums.sort()
        self.genCntSums()
        print(self.numToCntDict)
        print(self.nums)
        print(self.cntSums)

    def genNumToCntDict(self, nums):
        for num in nums:
            cnt = self.numToCntDict.get(num)
            if cnt is None:
                self.numToCntDict[num] = 0
            self.numToCntDict[num] += 1

    def biSearch(self, target, lower, upper):
        while lower < upper:
            middle = (lower + upper) / 2
            if target > self.nums[middle] and target <= self.nums[middle + 1]:
                return middle
            elif target <= self.nums[middle]:
                upper = middle
            elif target > self.nums[middle + 1]:
                lower = middle + 1
        return lower if target > self.nums[lower] else -1

    def genCntSums(self):
        numsLen = len(self.nums)
        self.cntSums = [0] * numsLen
        self.cntSums[0] = self.numToCntDict[self.nums[0]]
        for i, num in enumerate(self.nums[1:]):
            self.cntSums[i+1] = self.cntSums[i] + self.numToCntDict[self.nums[i+1]]

if __name__ == '__main__':
    solution = Solution()
    nums = [2,4,5,7,3,2,4,6,7,2]
    solution.triangleNumber(nums)
