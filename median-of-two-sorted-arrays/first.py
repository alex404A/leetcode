class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        len1 = len(nums1); len2 = len(nums2)
        medianIndex = (len1 + len2) / 2
        if (len1 + len2) % 2 == 1:
            return self.getNthNum(nums1, nums2, medianIndex + 1)
        else:
            return (self.getNthNum(nums1, nums2, medianIndex) + self.getNthNum(nums1, nums2, medianIndex + 1)) / 2.0

    def getNthNum(self, nums1, nums2, n):
        shorterNums = nums1 if len(nums1) <= len(nums2) else nums2
        longerNums = nums1 if len(nums1) > len(nums2) else nums2
        if len(shorterNums) == 0: return longerNums[n - 1]
        if (n == 1): return  min(shorterNums[0], longerNums[0])
        index = min(n / 2, len(shorterNums))
        if shorterNums[index - 1] <= longerNums[index - 1]:
            return self.getNthNum(shorterNums[index:], longerNums, n - index)
        else:
            return self.getNthNum(shorterNums, longerNums[index:], n - index)



if __name__ == '__main__':
    solution = Solution()
    nums1 = [0, 2, 4, 6, 8]
    nums2 = [1, 3, 5, 7, 9]
    print(solution.findMedianSortedArrays(nums1, nums2))
