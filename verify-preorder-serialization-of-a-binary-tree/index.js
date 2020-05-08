var isValidSerialization = function(preorder) {
  var valCnt = 0
  var nullCnt = 0
  var items = preorder.split(',')

  console.log(items)

  for (var i = 0; i < items.length; i++) {
    var c = items[i]
    if (c === '#') {
      nullCnt++
    } else {
      valCnt++
    }
    if (valCnt + 1 <= nullCnt && i !== items.length - 1) {
      console.log(valCnt)
      console.log(nullCnt)
      return false
    }
  }
  return valCnt + 1 === nullCnt
};

var preorder = '9,3,4,#,#,1,#,#,2,#,6,#,#'
console.log(isValidSerialization(preorder))
