class Solution(object):
    def spiralOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        if len(matrix) == 0: return []
        rowNum = len(matrix)
        colNum = len(matrix[0])
        leftUpperPoint = [0, 0]
        rightUpperPoint = [0, colNum - 1]
        rightLowerPoint = [rowNum - 1, colNum - 1]
        leftLowerPoint = [rowNum - 1, 0]
        results = []
        while rightUpperPoint[1] - leftUpperPoint[1] >= 0 and leftLowerPoint[0] - leftUpperPoint[0] >= 0:
            for i in range(leftUpperPoint[1], rightUpperPoint[1] + 1):
                results.append(matrix[leftUpperPoint[0]][i])

            for i in range(rightUpperPoint[0] + 1, rightLowerPoint[0]):
                results.append(matrix[i][rightUpperPoint[1]])

            if leftLowerPoint[0] != leftUpperPoint[0]:
                for i in range(rightLowerPoint[1], leftLowerPoint[1] - 1, -1):
                    results.append(matrix[rightLowerPoint[0]][i])

            if rightUpperPoint[1] != leftUpperPoint[1]:
                for i in range(leftLowerPoint[0] - 1, leftUpperPoint[0], -1):
                    results.append(matrix[i][leftLowerPoint[1]])
            leftUpperPoint[0] += 1
            leftUpperPoint[1] += 1
            rightUpperPoint[0] += 1
            rightUpperPoint[1] -= 1
            rightLowerPoint[0] -= 1
            rightLowerPoint[1] -= 1
            leftLowerPoint[0] -= 1
            leftLowerPoint[1] += 1

        return results

if __name__ == '__main__':
    solution = Solution()
    matrix = [[7], [9], [6]]
    print(solution.spiralOrder(matrix))
