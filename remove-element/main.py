class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        length = len(nums)
        if length == 0:
            return 0;
        i, j = 0, 0
        while i < length and i + j < length - 1:
            while nums[i] == val and i + j < length - 1:
                nums[i], nums[length-j-1] = nums[length-j-1], nums[i]
                j += 1
            i += 1
        return length-j if nums[length-j-1] != val else length-j-1

if __name__ == '__main__':
    solution = Solution()
    nums1 = [3, 2, 2, 3]
    nums2 = [3, 3, 3]
    length1 = solution.removeElement(nums1, 3)
    length2 = solution.removeElement(nums2, 3)
    print(length1, nums1)
    print(length2, nums2)
