class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 0:
            return 0
        boundarySum = nums[0]
        maxSubSum = nums[0]
        for num in nums[1:]:
            if boundarySum >= 0:
                boundarySum += num
            else:
                boundarySum = num
            if boundarySum >= maxSubSum:
                maxSubSum = boundarySum
        return maxSubSum

if __name__ == '__main__':
    solution = Solution()
    nums1 = [-2,1,-3,4,-1,2,1,-5,4]
    nums2 = [-2,-1]
    print(solution.maxSubArray(nums1))
    print(solution.maxSubArray(nums2))
