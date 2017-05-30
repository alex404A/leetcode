class Solution(object):
    def __init__(self):
        self.results = [1, 1]

    def genValidCountList(self, num = 30):
        print(num)
        self.results += [0] * (num - 1)
        self.getValidCount(num)
        return self.results[:-2]

    def getValidCount(self, num):
        result = 0
        for i in range(num - 2, -1, -1):
            if self.results[i] != 0:
                result += self.results[i]
            else:
                result += self.getValidCount(i)
        self.results[num] = result
        return result

if __name__ == '__main__':
    solution = Solution()
    print(solution.genValidCountList(32))
