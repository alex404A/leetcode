class Solution(object):

    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """

        left = 0
        right = len(nums) - 1
        while left <= right:
            middle = (left + right) / 2
            if nums[middle] == target:
                return middle
            if nums[middle] >= nums[left]:
                if target < nums[middle] and target >= nums[left]:
                    right = middle - 1
                else:
                    left = middle + 1
            elif nums[middle] <= nums[right]:
                if target > nums[middle] and target <= nums[right]:
                    left = middle + 1
                else:
                    right = middle - 1
        return -1

if __name__ == '__main__':
    solution = Solution()
    nums1 = [4,5,6,7,8,9,1,2,3]
    print(solution.search(nums1, 1))
    print(solution.search(nums1, 3))
