class Solution(object):
    def singleNonDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        def getResult(nums):
            print(nums)
            results = {}
            for num in nums:
                if results.get(num) is None:
                    results[num] = 0
                results[num] += 1
            for num, cnt in results.iteritems():
                if cnt == 1:
                    return num

        if len(nums) <= 5:
            return getResult(nums)
        cutIndex = len(nums) >> 1
        firstParts = nums[:cutIndex]
        secondParts = nums[cutIndex:]
        if len(firstParts) % 2 == 0:
            if firstParts[-1] == firstParts[-2]:
                return self.singleNonDuplicate(secondParts)
            else:
                return self.singleNonDuplicate(firstParts[:-1])
        else:
            if firstParts[-1] == firstParts[-2]:
                return self.singleNonDuplicate(firstParts[:-2])
            else:
                if firstParts[-1] == secondParts[0]:
                    return self.singleNonDuplicate(secondParts[1:])
                else:
                    return firstParts[-1]

if __name__ == '__main__':
    solution = Solution()
    nums = [1,1,2]
    print(solution.singleNonDuplicate(nums))
    nums = [1,1,2,2,3,3,4,4,5]
    print(solution.singleNonDuplicate(nums))
    nums = [1,1,2,3,3,4,4]
    print(solution.singleNonDuplicate(nums))
    nums = [0,1,1,2,2,3,3,4,4,5,5]
    print(solution.singleNonDuplicate(nums))
