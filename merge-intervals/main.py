# Definition for an interval.
class Interval(object):
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        nums = []
        for interval in intervals:
            start = interval.start
            end = interval.end
            startIndex = self.biSearch(nums, start)
            endIndex = self.biSearch(nums, end)
            isStartInInterval = startIndex % 2
            isEndInInterval = endIndex % 2
            if isStartInInterval and isEndInInterval:
                nums = nums[0:startIndex] + nums[endIndex:]
            elif not isStartInInterval and isEndInInterval:
                nums = nums[0:startIndex] + [start] + nums[endIndex:]
            elif isStartInInterval and not isEndInInterval:
                nums = nums[0:startIndex] + [end] + nums[endIndex:]
            else:
                nums = nums[0:startIndex] + [start, end] + nums[endIndex:]

        results = []
        nonDuplicateNums = []
        for i in range(len(nums)):
            ndLen = len(nonDuplicateNums)
            if ndLen != 0 and i % 2 == 0 and nonDuplicateNums[ndLen - 1] == nums[i]:
                nonDuplicateNums.pop()
            else:
                nonDuplicateNums.append(nums[i])
            print(nonDuplicateNums)
        for i in range(len(nonDuplicateNums) / 2):
            interval = Interval(nonDuplicateNums[i*2], nonDuplicateNums[i*2 + 1])
            results.append(interval)
        return results

    def biSearch(self, array, target):
        def search(lower, upper):
            upper = upper if upper != len(array) else len(array) - 1
            for i in range(upper, lower, -1):
                if target <= array[i] and target > array[i - 1]:
                    return i
            return lower if target <= array[lower] else upper + 1
        lower = 0
        upper = len(array)
        while lower < upper:
            if upper - lower <= 2:
                return search(lower, upper)
            x = lower + (upper - lower) / 2
            val = array[x]
            if target == val:
                return x
            elif target > val:
                if lower == x:
                    break
                lower = x
            elif target < val:
                if upper == x:
                    break
                upper = x
        return 0

    def printInterval(self, intervals):
        nums = []
        for interval in intervals:
            nums.append(interval.start)
            nums.append(interval.end)
        print(nums)

if __name__ == '__main__':
    solution = Solution()
    nums = [1, 4, 0, 0]
    intervals = []
    for i in range(len(nums) / 2):
        interval = Interval(nums[i*2], nums[i*2 + 1])
        intervals.append(interval)
    solution.printInterval(intervals)
    intervals = solution.merge(intervals)
    solution.printInterval(intervals)
