/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function(coins, amount) {
  coins.sort((a, b) => a > b ? 1 : a === b ? 0 : -1)
  var mem = {}
  return check(coins, coins.length - 1, amount, mem)
}

var max = Number.MAX_SAFE_INTEGER;

function check(coins, index, amount, mem) {
  if (amount === 0) {
    return 0
  }
  if (amount < 0) {
    return -1
  }
  if (mem[amount + '_' + index] !== undefined) {
    return mem[amount + '_' + index]
  }
  var result = max
  for (var i = index; i >= 0; i--) {
    var coin = coins[i]
    var cur = check(coins, i, amount - coin, mem)
    if (cur >= 0) {
      if (result > cur + 1) {
        result = cur + 1
      }
    }
  }
  mem[amount + '_' + index] = result === max ? -1 : result
  return mem[amount + '_' + index]
}

function test(coins, amount, expected) {
  var actual = coinChange(coins, amount)
  if (expected !== actual) {
    console.log(`coins ${coins} of amount ${amount} expected ${expected} actual ${actual}`)
  }
}

function main() {
  test([1,4,5], 12, 3)
  test([186,419,83,408], 6249, 20)
}

main()
