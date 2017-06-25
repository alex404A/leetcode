class Solution(object):
    def largestDivisibleSubset(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        if len(nums) == 0:
            return []

        nums.sort()
        subsetDict = {1: [[nums[0]]]}
        diffLenList = [1]
        for num in nums[1:]:
            isSubsetSearchedInAllLevel = False
            for i in range(0, len(diffLenList)):
                subsetList = subsetDict.get(diffLenList[i])
                isSubsetSearchedInLenLevel = False
                for subset in subsetList:
                    if num % subset[len(subset) - 1] == 0:
                        newSubset = subset + [num]
                        newSubsetLen = len(newSubset)
                        if subsetDict.get(newSubsetLen) is None:
                            subsetDict[newSubsetLen] = []
                            diffLenList = diffLenList[0:i] + [newSubsetLen] + diffLenList[i:]
                        subsetDict[newSubsetLen].append(newSubset)
                        isSubsetSearchedInLenLevel = True
                        isSubsetSearchedInAllLevel = True
                        break
                if isSubsetSearchedInLenLevel:
                    break
            if not isSubsetSearchedInAllLevel:
                subsetDict[1].append([num])

        return subsetDict[diffLenList[0]][0]

if __name__ == '__main__':
    solution = Solution()
    nums = [3,4,16,8]
    print(solution.largestDivisibleSubset(nums))
