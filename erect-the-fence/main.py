# Definition for a point.
class Point(object):
    def __init__(self, a=0, b=0):
        self.x = a
        self.y = b

class Solution(object):
    def __init__(self):
        self.points = []
        self.poles = {}
        self.intervals = {
            'upLeft': [],
            'upRight': [],
            'downLeft': [],
            'downRight': [],
        }

    def outerTrees(self, points):
        """
        :type points: List[Point]
        :rtype: List[Point]
        """
        def posCmp(slope1, slope2):
            return 1 if slope1 > slope2 else 0 if slope1 == slope2 else -1
        def negCmp(slope1, slope2):
            return -1 if slope1 > slope2 else 0 if slope1 == slope2 else 1
        self.points = points
        self.collectPoles()
        self.distPoints()
        results = []
        results += self.collectRopePoints(self.poles['left'], self.poles['up'], self.intervals['upLeft'], posCmp)
        results += self.collectRopePoints(self.poles['up'], self.poles['right'], self.intervals['upRight'], negCmp)
        results += self.collectRopePoints(self.poles['right'], self.poles['down'], self.intervals['downRight'], posCmp)
        results += self.collectRopePoints(self.poles['down'], self.poles['left'], self.intervals['downLeft'], negCmp)
        results += [pole for pole in self.poles.itervalues()]
        results = self.collectNonDuplicatePoints(results)
        self.printPoints('results', results)
        return results

    def collectNonDuplicatePoints(self, points):
        coordinateList = [(point.x, point.y) for point in points]
        return [point for point in dict(zip(coordinateList, points)).itervalues()]

    def collectPoles(self):
        xMax = xMin = self.points[0].x
        yMax = yMin = self.points[0].y
        poles = {'left': [], 'up': [], 'right': [], 'down': []}
        for index, point in enumerate(self.points):
            if point.x > xMax:
                xMax = point.x
                del poles['right'][:]
                poles['right'].append(point)
            elif point.x == xMax:
                poles['right'].append(point)
            if point.x < xMin:
                xMin = point.x
                del poles['left'][:]
                poles['left'].append(point)
            elif point.x == xMin:
                poles['left'].append(point)
            if point.y > yMax:
                yMax = point.y
                del poles['up'][:]
                poles['up'].append(point)
            elif point.y == yMax:
                poles['up'].append(point)
            if point.y < yMin:
                yMin = point.y
                del poles['down'][:]
                poles['down'].append(point)
            elif point.y == yMin:
                poles['down'] = point
        return poles
        # self.printPoints('poles', [self.poles['left']] + [self.poles['up']] + [self.poles['right']] + [self.poles['down']])

    def isFourPolesExisting(self, poles):
        if len(set(poles['left'] + poles['right'] + poles['up'] + poles['down'])) <= 3:
            return True

    def isPointsEqual(self, point1, point2):
        return True if point1.x == point2.x and point1.y == point2.y else False

    def distPoints(self):
        basicUpLeftSlope = self.calSlope(self.poles['up'], self.poles['left'])
        basicUpRightSlope = self.calSlope(self.poles['right'], self.poles['up'])
        basicDownRightSlope = self.calSlope(self.poles['down'], self.poles['right'])
        basicDownLeftSlope = self.calSlope(self.poles['left'], self.poles['down'])
        for point in self.points:
            if self.poles['left'] == point or self.poles['right'] == point or self.poles['up'] == point or self.poles['down'] == point:
                continue
            if point.x >= self.poles['left'].x and point.y >= self.poles['left'].y and point.x <= self.poles['up'].x and point.y <= self.poles['up'].y:
                currentSlope = self.calSlope(point, self.poles['left'])
                if currentSlope >= basicUpLeftSlope:
                    self.intervals['upLeft'].append(point)
            if point.x >= self.poles['up'].x and point.y <= self.poles['up'].y and point.x <= self.poles['right'].x and point.y >= self.poles['right'].y:
                currentSlope = self.calSlope(point, self.poles['up'])
                if currentSlope <= basicUpRightSlope:
                    self.intervals['upRight'].append(point)
            if point.x <= self.poles['right'].x and point.y <= self.poles['right'].y and point.x >= self.poles['down'].x and point.y >= self.poles['down'].y:
                currentSlope = self.calSlope(point, self.poles['right'])
                if currentSlope >= basicDownRightSlope:
                    self.intervals['downRight'].append(point)
            if point.x <= self.poles['down'].x and point.y >= self.poles['down'].y and point.x >= self.poles['left'].x and point.y <= self.poles['left'].y:
                currentSlope = self.calSlope(point, self.poles['down'])
                if currentSlope <= basicDownLeftSlope:
                    self.intervals['downLeft'].append(point)
        self.intervals['upLeft'].sort(key = lambda item: (item.x, item.y))
        self.intervals['upRight'].sort(key = lambda item: (item.x, 0 - item.y))
        self.intervals['downRight'].sort(key = lambda item: (0 - item.x, 0 - item.y))
        self.intervals['downLeft'].sort(key = lambda item: (0 - item.x, item.y))
        self.printPoints('dist-up-left', self.intervals['upLeft'])
        self.printPoints('dist-up-right', self.intervals['upRight'])
        self.printPoints('dist-down-right', self.intervals['downRight'])
        self.printPoints('dist-down-left', self.intervals['downLeft'])

    def collectRopePoints(self, startPole, endPole, rawPoints, compFunc):
        results = []
        maxSlopePoints = []
        points = rawPoints
        startPoint = startPole
        basicSlope = self.calSlope(startPoint, endPole)
        maxSlope = basicSlope
        lastMaxSlopeIndex = 0
        while len(points) > 0:
            for index, point in enumerate(points):
                currentSlope = self.calSlope(startPoint, point)
                compResult = compFunc(currentSlope, maxSlope)
                if compResult == 1:
                    del maxSlopePoints[:]
                    maxSlope = currentSlope
                if compResult >= 0:
                    lastMaxSlopeIndex = index
                    maxSlopePoints.append(point)
            results += maxSlopePoints
            if compFunc(maxSlope, basicSlope) == 0:
                break
            startPoint = maxSlopePoints[len(maxSlopePoints) - 1]
            del maxSlopePoints[:]
            basicSlope = self.calSlope(startPoint, endPole)
            points = points[lastMaxSlopeIndex + 1:]
            maxSlope = basicSlope
        return results

    def calSlope(self, point1, point2):
        if point1.x == point2.x:
            return float('infinity')
        return abs((point1.y - point2.y + 0.0) / (point1.x - point2.x))

    def printPoints(self, title, points):
        print(title)
        for point in points:
            print('(' + str(point.x) + ', ' + str(point.y) + ')')

if __name__ == '__main__':
    solution = Solution()
    points2 = [
        Point(3, 3), Point(9, 3), Point(4, 7), Point(9, 9), Point(8, 7),
        Point(4, 1), Point(0, 3), Point(2, 7)
    ]
    solution.outerTrees(points2)
