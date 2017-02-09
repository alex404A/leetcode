/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
  if (x === 0) {
    return 0;
  }
  var str = x.toString();
  var neg = false;
  if (str.startsWith('-')) {
    x = -x;
    str  = str.slice(1);
    neg = true;
  }

  var result = '';
  var len = str.length;
  var divisor, quotient;
  var remainder = x;
  for (var i = len-1; i >= 0; i--) {
    divisor = Math.pow(10, i);
    quotient = parseInt(remainder/divisor);
    remainder = remainder%divisor;
    result = quotient.toString() + result;
  }

  result = parseInt(result);
  if (result > Math.pow(2, 31)) {
    return 0;
  } else {
    if (neg) {
      return -result;
    } else {
      return result;
    }
  }

};

console.log(reverse(123));
console.log(reverse(-123));
console.log(reverse(0));
