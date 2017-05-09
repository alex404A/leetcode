class Solution(object):
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        numsLen = len(nums)
        if numsLen <= 1:
            return True
        maxStepIndex = 0
        maxDistance = nums[0]
        while True:
            print(maxStepIndex)
            if nums[maxStepIndex] == 0:
                return False
            if nums[maxStepIndex] >= numsLen - maxStepIndex - 1:
                return True
            for i in range(maxStepIndex + 1, maxStepIndex + nums[maxStepIndex] + 1):
                distance = i + nums[i]
                if distance >= maxDistance:
                    maxStepIndex = i
                    maxDistance = distance

if __name__ == '__main__':
    solution = Solution()
    nums = [1, 1, 2, 2, 0, 1, 1]
    print(solution.canJump(nums))
