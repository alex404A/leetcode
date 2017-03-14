class Solution(object):
    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """

        horizontalDict = {}
        verticalDist = {}
        squareDict = {}

        for listIndex, li in enumerate(board):
            for itemIndex, item in enumerate(li):
                if item == '.':
                    continue
                horizontalChildDict = horizontalDict.get(listIndex)
                if horizontalChildDict is None:
                    horizontalChildDict = {}
                    horizontalDict[listIndex] = horizontalChildDict
                else:
                    if horizontalChildDict.get(item) is not None:
                        print('horizotal', listIndex, itemIndex, horizontalChildDict)
                        return False
                horizontalChildDict[item] = 0
                verticalChildDist = verticalDist.get(itemIndex)
                if verticalChildDist is None:
                    verticalChildDist = {}
                    verticalDist[itemIndex] = verticalChildDist
                else:
                    if verticalChildDist.get(item) is not None:
                        print('vertical', listIndex, itemIndex, verticalChildDist)
                        return False
                verticalChildDist[item] = 0
                squareDictIndex = (listIndex / 3) * 3 + itemIndex / 3
                squareChildDist = squareDict.get(squareDictIndex)
                if squareChildDist is None:
                    squareChildDist = {}
                    squareDict[squareDictIndex] = squareChildDist
                else:
                    if squareChildDist.get(item) is not None:
                        print('square', listIndex, itemIndex, squareChildDist)
                        return False
                squareChildDist[item] = 0

        return True

if __name__ == '__main__':
    solution = Solution()
    list1 = ['.', '8', '7', '6', '5', '4', '3', '2', '1']
    list2 = ['2', '.', '.', '.', '.', '.', '.', '.', '.']
    list3 = ['3', '.', '.', '.', '.', '.', '.', '.', '.']
    list4 = ['4', '.', '.', '.', '.', '.', '.', '.', '.']
    list5 = ['5', '.', '.', '.', '.', '.', '.', '.', '.']
    list6 = ['6', '.', '.', '.', '.', '.', '.', '.', '.']
    list7 = ['7', '.', '.', '.', '.', '.', '.', '.', '.']
    list8 = ['8', '.', '.', '.', '.', '.', '.', '.', '.']
    list9 = ['9', '.', '.', '.', '.', '.', '.', '.', '.']
    board = [list1, list2, list3, list4, list5, list6, list7, list8, list9]
    print(solution.isValidSudoku(board))
