# Definition for a point.
class Point(object):
    def __init__(self, a=0, b=0):
        self.x = a
        self.y = b

class Solution(object):

    def outerTrees(self, points):
        """
        :type points: List[Point]
        :rtype: List[Point]
        """
        result = set()
        firstIndex, firstPoint = self.findMostLeftPoint(points)
        self.printPoints('left', [firstPoint])
        result.add(firstPoint)
        curIndex = firstIndex
        curPoint = firstPoint
        while True:
            nextPoint = points[0]
            nextIndex = 0
            colinearPoints = []
            for i, point in enumerate(points[1:]):
                if i == curIndex:
                    continue
                slope = self.checkSlope(curPoint, nextPoint, point)
                if slope > 0:
                    del colinearPoints[:]
                    nextPoint = point
                    nextIndex = i
                if slope == 0:
                    isColinearPointFurther = (self.getDistanceSquare(curPoint, point) - self.getDistanceSquare(curPoint, nextPoint)) > 0
                    if isColinearPointFurther is True:
                        colinearPoints.append(nextPoint)
                        nextPoint = point
                        nextIndex = i
                    else:
                        colinearPoints.append(point)
            self.printPoints('next', [nextPoint])
            result.add(nextPoint)
            for point in colinearPoints:
                result.add(point)
            if firstIndex == nextIndex:
                break
            curIndex = nextIndex
            curPoint = nextPoint
        self.printPoints('result', list(result))
        return list(result)

    def findMostLeftPoint(self, points):
        first = points[0]
        index = 0
        xMin = first.x
        for i, point in enumerate(points):
            if point.x < xMin:
                first = point
                index = 0
        return (index, first)

    # p1: cur, p2: next, p3: point
    def checkSlope(self, p1, p2, p3):
        return (p3.y - p2.y) * (p3.x - p1.x) - (p3.y - p1.y) * (p3.x - p2.x)

    def getDistanceSquare(self, p1, p2):
        return pow((p2.y - p1.y), 2) + pow((p2.x - p1.x), 2)

    def printPoints(self, title, points):
        print(title)
        for point in points:
            print('(' + str(point.x) + ', ' + str(point.y) + ')')

if __name__ == '__main__':
    points = [
        Point(0, 8), Point(9, 8), Point(2, 4)
    ]
    solution = Solution()
    solution.outerTrees(points)
