class Solution(object):
    def rotate(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: void Do not return anything, modify matrix in-place instead.
        """
        length = len(matrix)
        rotatedMatrix = []
        for i in range(length):
            rotatedMatrix.append([0] * length)
        for i in range(length):
            for j in range(length):
                rotatedMatrix[j][length - 1 - i] = matrix[i][j]
        for i in range(length):
            for j in range(length):
                matrix[i][j] = rotatedMatrix[i][j]
        print(matrix)

if __name__ == '__main__':
    solution = Solution()
    matrix = [[1,2,3],[4,5,6],[7,8,9]]
    solution.rotate(matrix)
