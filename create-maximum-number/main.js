/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @param {number} k
 * @return {number[]}
 */
var maxNumber = function(nums1, nums2, k) {
  var i = j = 0
  var p1
  var p2
  var nums
  var result = ['0']
  for (; i <= nums1.length; i++) {
    j = k - i
    if (j >= 0 && j <= nums2.length) {
      p1 = findMax(nums1, i)
      p2 = findMax(nums2, j)
      nums = merge(p1, p2)
      result = getMax(nums, result)
      console.log("nums1: " + nums1 + " length: " + i)
      console.log("nums2: " + nums2 + " length: " + j)
      console.log("merge: " + nums)
    }
  }
  return result
};

function findMax(nums, k) {
  var stack = []
  for (var i = 0; i < nums.length; i++) {
    while (stack.length > 0 && nums.length - i > k - stack.length) {
      if (stack[stack.length - 1] < nums[i]) {
        stack.pop()
      } else {
        break
      }
    }
    stack.push(nums[i])
  }
  return stack.slice(0, k)
}

function merge(nums1, nums2) {
  var result = []
  var temp = []
  while (nums1.length > 0 && nums2.length > 0) {
    if (nums1.join('') > nums2.join('')) {
      result.push(nums1.shift())
    } else {
      result.push(nums2.shift())
    }
  }
  if (nums1.length > 0) {
    return result.concat(nums1)
  } else {
    return result.concat(nums2)
  }
}

function getMax(newNums, oldNums) {
  var newResult = newNums.join('')
  var oldResult = oldNums.join('')
  return newResult > oldResult ? newNums : oldNums
}

nums1 = [6, 7]
nums2 = [6, 7]
k = 4 
console.log(maxNumber(nums1, nums2, k))
