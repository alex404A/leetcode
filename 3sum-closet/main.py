class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        mindiff = float('inf')
        result = 0
        nums.sort()
        for i in range(len(nums)):
            left = i + 1
            right = len(nums) - 1
            while left < right:
                tmp = nums[i] + nums[left] + nums[right]
                diff = abs(tmp - target)
                if diff < mindiff:
                    mindiff = diff
                    result = tmp
                if tmp == target:
                    return target
                elif tmp < target:
                    left += 1
                else:
                    right -= 1
        return result

if __name__ == '__main__':
    solution = Solution()
    nums = [-1, 2, 1, -4]
    target = 1
    print(solution.threeSumClosest(nums, target))
