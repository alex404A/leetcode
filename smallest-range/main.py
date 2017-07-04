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
            if startIndex == -1:
                startIndex = 0
            if endIndex == -1:
                endIndex = len(nums) - 1
            intervalList.append([startIndex, endIndex])
            if nums[startIndex] < start:
                start = nums[startIndex]
            if nums[endIndex] > end:
                end = nums[endIndex]
        comb = (start, maxStart) if maxStart - start <= end - minEnd else (minEnd, end)
        print(intervalList)

        commonNumDict = {}
        for i, interval in enumerate(intervalList):
            nums = numsList[i]
            startIndex = interval[0]
            endIndex = interval[1]
            maxIndex = max(startIndex, endIndex)
            minIndex = min(startIndex, endIndex)
            for j in range(minIndex, maxIndex + 1):
                if commonNumDict.get(nums[j]) is None:
                    commonNumDict[nums[j]] = set()
                commonNumDict[nums[j]].add(i)

        validNums = [num for num in commonNumDict.iterkeys()]
        validNums.sort()
        print(validNums)
        for i in range(len(validNums) - 1):
            j = i + 1
            curNums = set(commonNumDict[validNums[i]])
            while j < len(validNums):
                if validNums[j] - validNums[i] >= comb[1] - comb[0]:
                    break
                for num in commonNumDict[validNums[j]]:
                    curNums.add(num)
                if len(curNums) == len(numsList):
                    comb = (validNums[i], validNums[j])
                    break
                j += 1

        return list(comb)


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
        [-38,15,17,18],
        [-34,46,58,59,61],
        [-55,-31,-13,64,82,82,83,84,85],
        [-3,63,70,90],
        [2,6,10,28,28,32,32,32,33],
        [-23,82,88,88,88,89],
        [33,60,72,74,75],
        [-5,44,44,57,58,58,60],
        [-29,-22,-4,-4,17,18,19,19,19,20],
        [22,57,82,89,93,94],
        [24,38,45],
        [-100,-56,41,49,50,53,53,54],
        [-76,-69,-66,-53,-27,-1,9,29,31,32,32,32,34],
        [22,47,56],
        [-34,-28,7,44]
    ]
    print(numsList)
    print(solutioon.smallestRange(numsList))
