class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        dict = {}
        for i in range(len(nums)):
            if target - nums[i] in dict:
                return sorted((i, dict[target - nums[i]]))
            dict[nums[i]] = i

if __name__ == "__main__":
    testList = [1, 2, 4, 7]
    solution = Solution()
    print solution.twoSum(testList, 6)
