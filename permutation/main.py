class Solution(object):
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        if len(nums) <= 1:
            return [nums]
        subNums = [[nums[0]]]
        for i in range(1, len(nums)):
            num = nums[i]
            subNums = self.permuteSubNums(num ,subNums)
        return subNums

    def permuteSubNums(self, num, subNums):
        results = []
        for i in range(len(subNums)):
            for j in range(len(subNums[i]) + 1):
                results.append(self.genNewNums(j, num, subNums[i]))
        return results

    def genNewNums(self, index, num, nums):
        firstParts = nums[0:index]
        secondParts = nums[index:]
        return firstParts + [num] + secondParts

if __name__ == '__main__':
    solution = Solution()
    nums = [1, 2, 3, 4, 5, 6]
    print(solution.permute(nums))
