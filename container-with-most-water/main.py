class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        result = 0
        left = 0
        right = len(height) - 1
        while left < right:
            currentHeight = min(height[left], height[right])
            currentWidth = right - left
            water = currentWidth * currentHeight
            if water > result:
                result = water
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return result
