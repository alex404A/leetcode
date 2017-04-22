class Solution(object):
    def __init__(self):
        self.cnt = 0

    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        remaindar = (len(nums1) + len(nums2)) % 2
        half = (len(nums1) + len(nums2)) / 2
        indexes = []
        if remaindar == 0:
            indexes = [half - 1, half]
        else:
            indexes =  [half]
        return self.findSpecificSortedArrays(nums1, nums2, indexes)

    def findSpecificSortedArrays(self, nums1, nums2, indexes):
        self.cnt += 1
        if self.cnt > 10:
            return 0
        print(indexes)
        print(nums1, nums2)
        if len(nums1) <= 2 or len(nums2) <= 2:
            def getResult(nums1, nums2, index):
                nums = nums1 if len(nums1) > 0 else nums2
                sum = 0
                for index in indexes:
                    sum += nums[index]
                return sum / len(nums)
            if len(nums1) == 0 or len(nums2) == 0:
                return getResult(nums1, nums2, index)
            else:
                # to be fixed
                return 0
        elif indexes[0] == 0:
            sum = 0
            for i in indexes:
                if len(nums1) == 0:
                    sum += nums2[0]
                    nums2 = nums2[1:]
                elif len(nums2) == 0:
                    sum += nums1[0]
                    nums1 = nums1[1:]
                else:
                    if nums1[0] <= nums2[0]:
                        sum += nums1[0]
                        nums1 = nums1[1:]
                    else:
                        sum += nums2[0]
                        nums2 = nums2[1:]
            return sum / len(indexes)
        shorterNums = nums1 if len(nums1) <= len(nums2) else nums2
        longerNums = nums1 if len(nums1) > len(nums2) else nums2
        medianIndex = (len(shorterNums) - 1) / 2
        medianInShorterNums = shorterNums[medianIndex]
        medianInLongerNums = longerNums[medianIndex]
        if (medianInLongerNums >= medianInShorterNums):
            shorterNums = shorterNums[medianIndex:]
        else:
            longerNums = longerNums[medianIndex:]
        indexes = [i - medianIndex for i in indexes]
        return self.findSpecificSortedArrays(shorterNums, longerNums, indexes)

if __name__ == '__main__':
    solution = Solution()
    nums1 = [1, 4, 6]
    nums2 = [2, 3, 5, 7, 8]
    print(solution.findMedianSortedArrays(nums1, nums2))
