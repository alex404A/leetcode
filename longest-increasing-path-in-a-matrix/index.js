/**
 * @param {number[][]} matrix
 * @return {number}
 */
var longestIncreasingPath = function(matrix) {
  if (matrix.length === 0 || matrix[0].length === 0) {
    return 0
  }
  var mem = initMem(matrix)
  var max = 1
  var m = matrix.length
  var n = matrix[0].length
  for (var i = 0; i < matrix.length; i++) {
    for (var j = 0; j < matrix[0].length; j++) {
      var len = dfs(i, j, matrix, mem, m, n)  
      max = Math.max(max, len)
    }
  }
  return max
}

function initMem(matrix) {
  var mem = []
  for (var i = 0; i < matrix.length; i++) {
    mem[i] = []
  }
  return mem
}

var dirs = [[0, 1], [1, 0], [0, -1], [-1, 0]]

function dfs(i, j, matrix, mem, m, n) {
  if (mem[i][j]) {
    return mem[i][j]
  }
  var max = 1
  for (var k = 0; k < dirs.length; k++) {
    var dir = dirs[k]
    x = i + dir[0]
    y = j + dir[1]
    if (x < 0 || x >= m || y < 0 || y >= n || matrix[i][j] >= matrix[x][y]) {
      continue
    }
    var len = 1 + dfs(x, y, matrix, mem, m, n)
    max = Math.max(len, max)
  }
  mem[i][j] = max
  return max
}

function buildKey(i, j) {
  return `${i}_${j}`
}

var matrix = [
  [3,4,5],
  [3,2,6],
  [2,2,7]
] 
console.log(longestIncreasingPath(matrix))
