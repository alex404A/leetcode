class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        len1 = len(nums1); len2 = len(nums2)
        medianIndex = (len1 + len2) / 2
        if (len1 + len2) % 2 == 1:
            return self.getNthNumForOdd(nums1, nums2, medianIndex + 1)
        else:
            return self.getNthNumForEven(nums1, nums2, medianIndex)

    def getNthNumForOdd(self, nums1, nums2, n):
        shorterNums = nums1 if len(nums1) <= len(nums2) else nums2
        longerNums = nums1 if len(nums1) > len(nums2) else nums2
        if len(shorterNums) == 0: return longerNums[n - 1]
        if (n == 1): return  min(shorterNums[0], longerNums[0])
        index = min(n / 2, len(shorterNums))
        if shorterNums[index - 1] <= longerNums[index - 1]:
            return self.getNthNumForOdd(shorterNums[index:], longerNums, n - index)
        else:
            return self.getNthNumForOdd(shorterNums, longerNums[index:], n - index)

    def getNthNumForEven(self, nums1, nums2, n):
        shorterNums = nums1 if len(nums1) <= len(nums2) else nums2
        longerNums = nums1 if len(nums1) > len(nums2) else nums2
        if len(shorterNums) == 0: return (longerNums[n - 1] + longerNums[n]) / 2.0
        if (n == 1):
            nums = shorterNums[0: 2] + longerNums[0: 2]
            nums.sort()
            return (nums[0] + nums[1]) / 2.0
        index = min(n / 2, len(shorterNums))
        if shorterNums[index - 1] <= longerNums[index - 1]:
            return self.getNthNumForEven(shorterNums[index:], longerNums, n - index)
        else:
            return self.getNthNumForEven(shorterNums, longerNums[index:], n - index)

if __name__ == '__main__':
    solution = Solution()
    nums1 = []
    nums2 = [2, 3]
    print(solution.findMedianSortedArrays(nums1, nums2))
