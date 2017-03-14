class Solution(object):
    # @param A, a list of integers
    # @param target, an integer to be searched
    # @return insert position
    def searchInsert(self, A, target):
        if len(A) == 0:
            return 0
        left = 0; right = len(A) - 1
        while left < right:
            mid = (left + right) / 2
            if A[mid] > target:
                right = mid - 1
            elif A[mid] < target:
                left = mid + 1
            else:
                return mid
        if target <= A[left]:
            return left
        else:
            return left + 1

if __name__ == '__main__':
    solution = Solution()
    nums1 = []
    nums2 = [1, 3, 5, 8]
    nums3 = [1, 3, 5, 8, 10]
    print(solution.searchInsert(nums1, 2))
    print(solution.searchInsert(nums2, 4))
    print(solution.searchInsert(nums2, 1))
    print(solution.searchInsert(nums3, 4))
    print(solution.searchInsert(nums3, 8))
