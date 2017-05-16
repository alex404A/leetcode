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
        self.points = points
        self.collectPoles()
        self.distPoints()
        self.printPoles()
        print(self.intervals)

    def collectPoles(self):
        xMax = xMin = self.points[0].x
        yMax = yMin = self.points[0].y
        self.poles['left'] = self.poles['right'] = points[0]
        self.poles['up'] = self.poles['down'] = points[0]
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

    def distPoints(self):
        for point in points:
            if self.poles['left'] == point or self.poles['right'] == point or self.poles['up'] == point or self.poles['down'] == point:
                continue
            if point.x >= self.poles['left'].x and point.y >= self.poles['left'].y and point.x < self.poles['up'].x and point.y < self.poles['up'].y:
                self.intervals['upLeft'].append({'point': point, 'k': (point.y - self.poles['left'].y + 0.0) / (point.x - self.poles['left'].x)})
            if point.x >= self.poles['up'].x and point.y <= self.poles['up'].y and point.x < self.poles['right'].x and point.y > self.poles['right'].y:
                self.intervals['upRight'].append({'point': point, 'k': (point.y - self.poles['up'].y + 0.0) / (point.x - self.poles['up'].x)})
            if point.x <= self.poles['right'].x and point.y <= self.poles['right'].y and point.x > self.poles['down'].x and point.y > self.poles['down'].y:
                self.intervals['downRight'].append({'point': point, 'k': (point.y - self.poles['right'].y + 0.0) / (point.x - self.poles['right'].x)})
            if point.x <= self.poles['down'].x and point.y >= self.poles['down'].y and point.x > self.poles['left'].x and point.y < self.poles['left'].y:
                self.intervals['downLeft'].append({'point': point, 'k': (point.y - self.poles['down'].y + 0.0) / (point.x - self.poles['down'].x)})
        self.intervals['upLeft'].sort(key = lambda item: item['point'].x)
        self.intervals['upRight'].sort(key = lambda item: item['point'].x)
        self.intervals['downLeft'].sort(key = lambda item: item['point'].x)
        self.intervals['downRight'].sort(key = lambda item: item['point'].x)

    def printPoles(self):
        print('poles:')
        for pos, point in self.poles.iteritems():
            print(pos + ': (' + str(point.x) + ', ' + str(point.y) + ')')

if __name__ == '__main__':
    solution = Solution()
    points = [Point(1, 1), Point(2, 2), Point(2, 0), Point(2, 4), Point(3, 3), Point(4, 2)]
    solution.outerTrees(points)
