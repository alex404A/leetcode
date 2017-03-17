class Solution(object):
    def __init__(self):
        self.board = []
        self.items = []
        self.searchedItemStack = []

    def solveSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """

        self.board = board
        scoreBoard = self.getScoreBoard()
        self.items = self.getItemTupleListOrderedByScore(scoreBoard)
        item = self.items.pop()
        self.tryOneItem(item)
        for li in self.board:
            print(li)

    def tryOneItem(self, item):
        print('step into', item)
        rowIndex, colIndex = item[0], item[1]
        rowItems = self.board[rowIndex]
        colItems = [li[colIndex] for li in self.board]
        squareItems = self.getSquareItems((rowIndex/3) * 3 + colIndex/3)
        items = rowItems + colItems + squareItems
        candidates = [i for i in ['1', '2', '3', '4', '5', '6', '7', '8', '9'] if i not in items]
        if (len(candidates) == 0):
            self.items.append(item)
            print('no candidate')
            return False
        self.searchedItemStack.append(item)
        if (len(self.items) == 0):
            self.board[rowIndex][colIndex] = candidates[0]
            return True
        for candidate in candidates:
            nextItem = self.items.pop()
            if (len(candidates) > 1):
                print('try candidate', candidate, item)
            self.board[rowIndex][colIndex] = candidate
            if self.tryOneItem(nextItem):
                return True
        self.board[rowIndex][colIndex] = '.'
        self.searchedItemStack.pop()
        self.items.append(item)
        return False

    def getItemTupleListOrderedByScore(self, scoreBoard):
        itemList = []
        rowScores = scoreBoard.get('row')
        columnScores = scoreBoard.get('column')
        squareScores = scoreBoard.get('square')
        for listIndex, li in enumerate(self.board):
            for colIndex, item in enumerate(li):
                score = rowScores[listIndex] + columnScores[colIndex] + squareScores[(listIndex/3) * 3 + colIndex/3]
                if item == '.':
                    score -= 3
                    itemList.append((listIndex, colIndex, score))
        return sorted(itemList, key=lambda item: item[2], reverse=True)

    def getScoreBoard(self):
        scoreBoard = {
            'row': [],
            'column': [],
            'square': []
        }
        for i in range(9):
            scoreBoard.get('row').append(self.getScore(self.board[i]))
            scoreBoard.get('column').append(self.getScore([li[i] for li in self.board]))
            scoreBoard.get('square').append(self.getScore(self.getSquareItems(i)))
        return scoreBoard

    def getScore(self, itemStr):
        score = 0
        for i in range(len(itemStr)):
            if itemStr[i] == '.':
                score += 1
        return score

    def getSquareItems(self, index):
        quotient = index / 3
        remainder = index % 3
        list0 = [li[remainder*3 + 0] for li in self.board[quotient*3: quotient*3 + 3]]
        list1 = [li[remainder*3 + 1] for li in self.board[quotient*3: quotient*3 + 3]]
        list2 = [li[remainder*3 + 2] for li in self.board[quotient*3: quotient*3 + 3]]
        return list0 + list1 + list2

if __name__ == '__main__':
    solution = Solution()
    list1 = ['5', '3', '.', '.', '7', '.', '.', '.', '.']
    list2 = ['6', '.', '.', '1', '9', '5', '.', '.', '.']
    list3 = ['.', '9', '8', '.', '.', '.', '.', '6', '.']
    list4 = ['8', '.', '.', '.', '6', '.', '.', '.', '3']
    list5 = ['4', '.', '.', '8', '.', '3', '.', '.', '1']
    list6 = ['7', '.', '.', '.', '2', '.', '.', '.', '6']
    list7 = ['.', '6', '.', '.', '.', '.', '2', '8', '.']
    list8 = ['.', '.', '.', '4', '1', '9', '.', '.', '5']
    list9 = ['.', '.', '.', '.', '8', '.', '.', '7', '9']
    board = [list1, list2, list3, list4, list5, list6, list7, list8, list9]
    solution.solveSudoku(board)
