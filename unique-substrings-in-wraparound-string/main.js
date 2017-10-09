/**
 *  * @param {string} p
 *   * @return {number}
 *    */
var findSubstringInWraproundString = function(p) {
  if (p.length === 0) return 0
  var possibleCmbs = getPossibleCmbs()
  var cntArr = getCntArr()

  function getIndex (l) {
    return l.charCodeAt(0) - 'a'.charCodeAt(0)
  }

  function getExtNum(startIndex, cnt, curIndex) {
    var remain
    if (startIndex - curIndex <= 0) {
      remain = curIndex - startIndex
    } else {
      remain = 26 - startIndex + curIndex
    }
    return cnt - remain
  }

  function iterToAdd(l) {
    var innerCnt = 1
    var i = getIndex(l)
    while (cnt - innerCnt >= 0) {
      if (innerCnt > cntArr[i]) {
        cntArr[i] = innerCnt
      }
      innerCnt += 1
      i = i === 0 ? 25 : i -1
    }
  }

  var start = p[0]
  var cnt = 1
  cntArr[getIndex(start)] = 1
  var index
  var extNum
  var j = 0
  for (var i = 1; i < p.length; i++) {
    if (possibleCmbs[p[i - 1] + p[i]]) {
      cnt += 1
      for (j = 0; j < cntArr.length; j++) {
        // console.log("startIndex: " + getIndex(start) + ", cnt: " + cnt + " curIndex: " + j)
        extNum = getExtNum(getIndex(start), cnt, j)
        // console.log(extNum)
        if (cntArr[j] < extNum) {
          cntArr[j] = extNum
        }
      }
    } else {
      start = p[i]
      cnt = 1
      index = getIndex(start)
      if (cntArr[index] === 0) {
        cntArr[index] = 1
      }
    }
    console.log(cntArr)
  }

  return cntArr.reduce(function (accu, cnt) {
    accu += cnt
    return accu
  }, 0)

};

function getPossibleCmbs () {
  var alphabets = 'abcdefghijklmnopqrstuvwxyz'
  var cmbs = {}
  for (var i = 0; i < alphabets.length - 1; i++) {
    cmbs[alphabets[i] + alphabets[i + 1]] = true
  }
  cmbs.za = true
  return cmbs
}

function getCntArr () {
  var result = []
  for (var i = 0; i < 26; i++) {
    result.push(0)
  }
  return result
}

var p = 'zabcdefghijklmnopqrstuvwxyz'
console.log(findSubstringInWraproundString(p))
