import re

class Excel(object):
    def __init__(self, H,W):
        """
        :type H: int
        :type W: str
        """
        self.vertexPattern = re.compile(r'(\w{2,}):(\w{2,})')
        self.sumCoordToDeps = {}
        self.letterToNum = {}
        self.matrix = []
        ordA = ord('A')
        orda = ord('a')
        for i in range(0, 26):
            self.letterToNum[chr(ordA + i)] = i
            self.letterToNum[chr(orda + i)] = i
        rowLen = self.letterToNum.get(W) + 1
        for i in range(0, H):
            self.matrix.append([0] * rowLen)

    def set(self, r, c, v):
        """
        :type r: int
        :type c: str
        :type v: int
        :rtype: void
        """
        coord = (r - 1, self.letterToNum[c])
        self.innerSet(coord, v)
        self.removeDeps(coord)

    def innerSet(self, coord, v):
        rawV = self.matrix[coord[0]][coord[1]]
        self.matrix[coord[0]][coord[1]] = v
        self.updateDepCoord(coord, rawV, v)

    def get(self, r, c):
        """
        :type r: int
        :type c: str
        :rtype: int
        """
        return self.matrix[r - 1][self.letterToNum[c]]

    def innerGet(self, vertex):
        return self.matrix[vertex[0]][vertex[1]]

    def sum(self, r, c, strs):
        """
        :type r: int
        :type c: str
        :type strs: List[str]
        :rtype: int
        """
        def getVertex(vertexStr):
            return (int(vertexStr[1:]) - 1, self.letterToNum[vertexStr[0]])

        def getVertices(strs):
            vertices = []
            for vertexStr in strs:
                match = self.vertexPattern.match(vertexStr)
                if match is None:
                    topLeft = getVertex(vertexStr)
                    vertices.append((topLeft, topLeft))
                else:
                    topLeft = getVertex(match.group(1))
                    btmRight = getVertex(match.group(2))
                    vertices.append((topLeft, btmRight))
            return vertices

        vertices = getVertices(strs)
        result = 0
        deps = {}
        for topLeft, btmRight in vertices:
            for i in range(topLeft[0], btmRight[0] + 1):
                for j in range(topLeft[1], btmRight[1] + 1):
                    result += self.matrix[i][j]
                    if deps.get((i, j)) is None:
                        deps[(i, j)] = 0
                    deps[(i, j)] += 1

        sumCoord = (r - 1, self.letterToNum[c])
        self.innerSet(sumCoord, result)
        self.removeDeps(sumCoord)
        self.addDeps(sumCoord, deps)
        return result

    def removeDeps(self, sumCoord):
        self.sumCoordToDeps.pop(sumCoord, None)

    def addDeps(self, sumCoord, deps):
        self.sumCoordToDeps[sumCoord] = deps

    def updateDepCoord(self, commonCoord, rawV, newV):
        for sumCoord, deps in self.sumCoordToDeps.iteritems():
            cnt = deps.get(commonCoord)
            if cnt is not None:
                rawSum = self.innerGet(sumCoord)
                self.innerSet(sumCoord, rawSum + (newV - rawV) * cnt)


if __name__ == '__main__':
    obj = Excel(26, 'Z')
    obj.set(1, 'A', 1)
    obj.set(1, 'I', 1)
    obj.sum(7 , 'D', ['A1:D6', 'A1:G3', 'A1:C12'])
    print(obj.sum(10 , 'G', ['A1:D7', 'D1:F10', 'D3:I8', 'I1:I9']))
