from typing import List

class Solution:

  def __init__(self):
    self.smaller = []

  def countSmaller(self, nums: List[int]) -> List[int]:
    self.smaller = [0] * len(nums)
    self.merge_sort(list(enumerate(nums)))
    return self.smaller
  
  def merge_sort(self, kv_list):
    mid = len(kv_list) // 2
    if mid == 0:
      return kv_list
    left = self.merge_sort(kv_list[0:mid])
    right = self.merge_sort(kv_list[mid:])
    m = len(left)
    n = len(right)
    i, j = 0, 0
    while i < m and j < n:
      if left[i][1] <= right[j][1]:
        kv_list[i+j] = left[i]
        self.smaller[left[i][0]] += j
        i += 1
      else:
        kv_list[i+j] = right[j]
        j += 1
    for k in range(i, m):
      self.smaller[left[k][0]] += n
      kv_list[k + j] = left[k]
    for k in range(j, n):
      kv_list[k + i] = right[k]
    return kv_list
    
if __name__ == '__main__':
  solution = Solution()
  nums = [5,2,2,7,8,6,1]
  smaller = solution.countSmaller(nums)
  print(smaller)
