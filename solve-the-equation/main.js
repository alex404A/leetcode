/**
 *  * @param {string} equation
 *   * @return {string}
 *    */
var solveEquation = function (equation) {
  equation += ' '
  left = parse(equation, 0)
  right = parse(equation, left.pos + 1)
  console.log(left)
  console.log(right)
  if (left.coefficient === right.coefficient) {
    return left.sum === right.sum ? 'Infinite solutions' : 'No solution'
  } else {
    return 'x=' + (right.sum - left.sum) / (left.coefficient - right.coefficient)
  }
};

function parse (str, start) {
  var reg = /[0-9]{1}/
  var wrapper = {
    sum: 0,
    coefficient: 0,
    pos: 0
  }
  var num = ''
  var l = ''
  var op = '+'
  for (var i = start; i < str.length; i++) {
    l = str[i]
    // console.log('item: ' + l + ' num: ' + num + ' operation: ' + op)
    // console.log(wrapper)
    if (reg.test(l)) {
      num += l 
    } else if (l === ' ' || l === '+' || l === '-' || l === '=') {
      if (num !== '') {
        wrapper.sum += cal(num, op)
        num = ''
      }
      if (l === '+' || l === '-') {
        op = l
      } else if (l === '=') {
        wrapper.pos = i
        break
      }
    } else if (l == 'x') {
      if (num === '') num = '1'
      wrapper.coefficient += cal(num, op)
      num = ''
    } else {
      throw 'Invalid item: ' + l
    }
  }
  return wrapper
}

function cal (num, op) {
  if (op === '-') {
    return 0 - parseInt(num)
  } else if (op === '+') {
    return parseInt(num)
  } else {
    throw 'Invalid Operation'
  }
}

var equation = "x+5-3+x=6+x-2"
console.log(equation)
console.log(solveEquation(equation))
