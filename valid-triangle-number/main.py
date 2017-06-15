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
        if (len(self.nums) == 0):
            return 0
        self.nums.sort()
        self.genCntSums()
        numsLen = len(self.nums)
        result = 0
        for j in range(numsLen):
            k = numsLen - 1
            secondCnt = self.numToCntDict[self.nums[j]]
            if secondCnt > 2:
                result += self.ncr(secondCnt, 3)
            leftStart = j - 1 if secondCnt == 1 else j
            for i in range(leftStart, -1, -1):
                rightStart = j + 1
                if secondCnt >= 2 and i < j:
                    rightStart = j
                if rightStart >= numsLen:
                    continue
                tmp = self.biSearch(self.nums[i] + self.nums[j], rightStart, k)
                if tmp == -1:
                    if i == j or rightStart == j:
                        continue
                    else:
                        break
                k = tmp
                thirdPart = self.cntSums[k] - self.cntSums[j]
                if rightStart == j:
                    result += self.ncr(secondCnt, 2) * self.numToCntDict[self.nums[i]]
                    result += self.numToCntDict[self.nums[i]] * self.numToCntDict[self.nums[j]] * thirdPart
                else:
                    if i == j:
                        result += self.ncr(secondCnt, 2) * thirdPart
                    else:
                        result += self.numToCntDict[self.nums[i]] * self.numToCntDict[self.nums[j]] * thirdPart
        lastCnt = self.numToCntDict[self.nums[numsLen - 1]]
        return result

    def genNumToCntDict(self, nums):
        for num in nums:
            if num == 0:
                continue
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
        self.cntSums = [0] * len(self.nums)
        self.cntSums[0] = self.numToCntDict[self.nums[0]]
        for i, num in enumerate(self.nums[1:]):
            self.cntSums[i+1] = self.cntSums[i] + self.numToCntDict[self.nums[i+1]]

    def ncr(self, n, r):
        small = min(r, n - r)
        if small == 0:
            return 1
        big = max(r, n - r)
        first = 1
        second = 1
        for i in range(2, small + 1):
            first *= i
        for i in range(big + 1, n + 1):
            second *= i
        return second / first

if __name__ == '__main__':
    solution = Solution()
    nums = [14,15,49,15,39,95,98,6,44]
    print(solution.triangleNumber(nums))
