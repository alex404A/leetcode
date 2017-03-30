import sys

class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        numsLen = len(nums)
        finalNums = [0] * numsLen
        for i, num in enumerate(nums):
            if num > 0 and num <= numsLen:
                finalNums[num - 1] = num
        for i, num in enumerate(finalNums):
            if num != i + 1:
                return i + 1
        return len(finalNums) + 1

if __name__ == '__main__':
    solution = Solution()
    nums = [5, 3, 9, 8, 2, 1, 4, 6, 7, 0]
    print(solution.firstMissingPositive(nums))
