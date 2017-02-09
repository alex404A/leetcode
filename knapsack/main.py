
class Solution(object):
    def knapsack(self, limit, vw):
        if len(vw) == 1:
            return vw[0][0] if vw[0][1] <= limit else 0
        if vw[-1][1] <= limit:
            withLastOne = self.knapsack(limit-vw[-1][1], vw[0:-1]) + vw[-1][0]
        else:
            withLastOne = 0
        withoutLastOne = self.knapsack(limit, vw[0:-1])
        return max(withLastOne, withoutLastOne)

if __name__ == '__main__':
    limit = 16
    vw = [(30, 4), (20, 5), (160, 16), (40, 10), (10, 3), (50, 9)]
    solution = Solution()
    print solution.knapsack(limit, vw)
