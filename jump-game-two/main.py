class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 1:
            return 0
            
        def cal(index):
            nextIndex = nums[index] + index
            if nextIndex >= len(nums) - 1:
                return True
            maxIndex = nextIndex
            for i in range(index + 1, nextIndex):
                if i + nums[i] > maxIndex + nums[maxIndex]:
                    maxIndex = i
            results.append(maxIndex)
            print(results)
            return False

        results = [0]
        while True:
            isBreak = cal(results[len(results) - 1])
            if isBreak:
                return len(results)


if __name__ == '__main__':
    solution = Solution()
    nums = [2, 3, 1, 1, 4, 3, 1, 8, 9, 10, 6]
    print(solution.jump(nums))
