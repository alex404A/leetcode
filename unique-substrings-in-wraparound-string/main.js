/**
 *  * @param {string} p
 *   * @return {number}
 *    */
var findSubstringInWraproundString = function(p) {
  if (p.length === 0) return 0
  var possibleCmbs = getPossibleCmbs()
  var map = {}
  map[p[0]] = true
  for (var i = 1; i < p.length; i++) {
  }
};

function getPossibleCmbs () {
  var alpabets = 'abcdefghijklmnopqrstuvwxyz'
  var cmbs = {}
  for (var i = 0; i < alphabets.length - 1; i++) {
    cmbs[alphabets[i] + alphabets[i + 1]] = true
  }
  return cmbs.za = true
}
