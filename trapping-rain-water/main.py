class Solution(object):
    def trap(self, heights):
        """
        :type height: List[int]
        :rtype: int
        """
        left = 0
        right = len(heights) - 1
        maxLeft = 0
        maxRight = 0
        result = 0
        while left < right:
            if heights[left] < heights[right]:
                if heights[left] >= maxLeft:
                    maxLeft = heights[left]
                else:
                    result += maxLeft - heights[left]
                left += 1
            else:
                if heights[right] >= maxRight:
                    maxRight = heights[right]
                else:
                    result += maxRight - heights[right]
                right -= 1
        return result

if __name__ == '__main__':
    solution = Solution()
    # height = [0,1,0,3,1,0,1,1,2,1,2,1]
    # height = [0,1,0,2,1,0,1,3,2,1,2,1]
    height = [5,5,4,7,8,2,6,9,4,5]
    # height = [4,3,3,9,3,0,9,2,8,3]
    print(solution.trap(height))
