/**
 *  * @param {string} num
 *   * @param {number} k
 *    * @return {string}
 *     */
var removeKdigits = function(num, k) {
  var stack = [] 
  var cnt = 0
  var len = 0
  var cur
  for (var i = 0; i < num.length; i++) {
    len = stack.length
    cur = num[i]
    console.log(cur)
    while (len > 0 && cnt < k) {
      if (stack[len - 1] > cur) {
        stack.pop()
        cnt += 1
        len = stack.length
      } else {
        break
      }
    }
    stack.push(cur)
    if (cnt === k) return finalProcess(stack.join('') + num.slice(i + 1))
  }
  return finalProcess(stack.slice(0, stack.length - (k - cnt)).join(''))
};

function finalProcess(newNum) {
  for (var i = 0; i < newNum.length; i++) {
    if (newNum[i] !== '0') break
  }
  return i !== newNum.length ? newNum.slice(i) : '0'
}

var num = '10'
console.log(removeKdigits(num, 2))
