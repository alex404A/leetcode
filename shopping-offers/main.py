class Solution(object):
  def __init__(self):
    self.statistics = {}

  def shoppingOffers(self, price, special, needs):
    """
    :type price: List[int]
    :type special: List[List[int]]
    :type needs: List[int]
    :rtype: int
    """
    return self.shopping(price, special, needs, 0)

  def shopping(self, price, special, needs, index):
    key = self.assembleTuple(needs, index)
    total = self.statistics.get(key)
    if total is not None:
      return total

    if (index == len(special)):
      total = self.shoppingWithOriginalPrice(price, needs)
      self.statistics[key] = total
      return total

    clone = list(needs)
    cnt = 0
    curOffer = special[index]
    for i in range(len(curOffer) - 1):
      if clone[i] - curOffer[i] >= 0:
        cnt += 1
        clone[i] -= curOffer[i]
      else:
        break;
    
    if cnt == len(curOffer) - 1:
      total = min(
        self.shopping(price, special, clone, index) + curOffer[cnt],
        self.shopping(price, special, needs, index + 1)
      )
    else:
      total = self.shopping(price, special, needs, index + 1)

    self.statistics[key] = total
    return total
  
  def shoppingWithOriginalPrice(self, price, needs):
    total = 0
    for i in range(len(needs)):
      total += needs[i] * price[i]
    return total
  
  def assembleTuple(self, needs, index):
    clone = list(needs)
    clone.append(index)
    return tuple(clone)

if __name__ == '__main__':
  solution = Solution()
  price = [2, 3, 4]
  special = [[1, 1, 0, 4], [2, 2, 1, 9]]
  needs = [1, 2, 1]
  print(solution.shoppingOffers(price, special, needs))
        