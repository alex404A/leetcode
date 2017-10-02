/**
 * @param {string} s
 * @return {string}
 */
var removeDuplicateLetters = function(s) {
    var cntMap = s.split('')
      .reduce(function (accu, letter) {
        if (!accu[letter]) {
          accu[letter] = 0
        }
        accu[letter] += 1
        return accu
      }, {})
    var stack = []
    var stackItemMap = {}
    var letter;
    var stackItem;
    for (var i = 0; i < s.length; i++) {
      letter = s[i]
      cntMap[letter] -= 1
      if (stackItemMap[letter]) {
        continue;
      }
      for (var j = stack.length - 1; j >= 0; j--) {
        stackItem = stack[stack.length - 1];
        if (stackItem > letter && cntMap[stackItem] > 0) {
          stack.pop();
          stackItemMap[stackItem] = false;
        }
      }
      stack.push(letter)
      stackItemMap[letter] = true
    }
    return stack.join('')
};

var s = 'cbacdcbc'
console.log(s)
console.log(removeDuplicateLetters(s))
