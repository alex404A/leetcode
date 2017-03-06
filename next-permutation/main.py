class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """

        partition = -1
        numsLen = len(nums)
        for i in range(numsLen-1, 0, -1):
            if nums[i] > nums[i-1]:
                partition = i - 1
                break

        if partition >= 0:
            for i in range(numsLen-1, partition, -1):
                if nums[i] > nums[partition]:
                    nums[i], nums[partition] = nums[partition], nums[i]
                    break

        for i in range(numsLen-1, partition, -1):
            nums[i], nums[partition + numsLen - i] = nums[partition + numsLen - i], nums[i]
            if (numsLen + partition) / 2 + 1 == i:
                break


if __name__ == '__main__':
    nums = [9, 6, 3]
    print('input: ' + str(nums))
    print('expected: [3, 6, 9]')
    solution = Solution()
    solution.nextPermutation(nums)
