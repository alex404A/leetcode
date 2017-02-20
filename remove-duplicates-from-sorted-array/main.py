class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if nums is None:
            return 0
        if len(nums) == 0:
            return 0
        j = 0;
        for i in range(0, len(nums)):
            if nums[i] != nums[j]:
                nums[i], nums[j+1] = nums[j+1], nums[i]
                j += 1
        return j + 1

if __name__ == '__main__':
    solution = Solution()
    nums = [1, 2, 2, 4]
    length = solution.removeDuplicates(nums)
    print(length, nums)
