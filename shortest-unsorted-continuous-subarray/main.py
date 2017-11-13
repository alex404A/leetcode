class Solution(object):
    def findUnsortedSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        if (len(nums) <= 1):
            return 0

        stack = []
        left = len(nums) - 1
        right = 0

        for i in range(len(nums)):
            while len(stack) > 0 and nums[stack[len(stack) - 1]] > nums[i]:
                last = stack.pop()
                left = min(left, last)
            stack.append(i)

        stack = []
        for i in range(len(nums) - 1, -1, -1):
            while len(stack) > 0 and nums[stack[len(stack) - 1]] < nums[i]:
                last = stack.pop()
                right = max(right, last)
            stack.append(i)
        
        return right - left + 1 if right >= left else 0


if __name__ == '__main__':
    solution = Solution()
    nums = [1, 2, 3, 4]
    solution.findUnsortedSubarray(nums)