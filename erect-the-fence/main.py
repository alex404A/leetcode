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
        self.poles['left'] = self.poles['right'] = self.points[0]
        self.poles['up'] = self.poles['down'] = self.points[0]
        for index, point in enumerate(self.points):
            if point.x > xMax:
                xMax = point.x
                self.poles['right'] = point
            if point.x < xMin:
                xMin = point.x
                self.poles['left'] = point
            if point.y > yMax:
                yMax = point.y
                self.poles['up'] = point
            if point.y < yMin:
                yMin = point.y
                self.poles['down'] = point
        self.printPoints('poles', [self.poles['left']] + [self.poles['up']] + [self.poles['right']] + [self.poles['down']])

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
        for key, value in self.intervals.iteritems():
            value.sort(key = lambda item: item.x)
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
                if  compResult == 1:
                    del maxSlopePoints[:]
                    maxSlope = currentSlope
                if compResult >= 0:
                    lastMaxSlopeIndex = index
                    maxSlopePoints.append(point)
            if maxSlope > basicSlope:
                results += maxSlopePoints
            elif maxSlope == basicSlope:
                results += maxSlopePoints
                break
            startPoint = maxSlopePoints[len(maxSlopePoints) - 1]
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
    points2 = [Point(3, 3), Point(1, 1), Point(2, 0), Point(2, 2), Point(4, 2), Point(2, 4)]
    solution.outerTrees(points2)
