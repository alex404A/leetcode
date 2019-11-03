class Solution(object):
    def canIWin(self, maxChoosableInteger, desiredTotal):
        """
        :type maxChoosableInteger: int
        :type desiredTotal: int
        :rtype: bool
        """
        if maxChoosableInteger * (maxChoosableInteger + 1) / 2 < desiredTotal:
            return False
        hmap = {}
        statusList = [0] * (maxChoosableInteger + 1)
        return self.judge(desiredTotal, statusList, hmap)

    def judge(self, desiredTotal, statusList, hmap):
        key = ''.join([str(i) for i in statusList])
        if hmap.get(key) is not None:
            return hmap.get(key)
        for i in range(1, len(statusList)):
            if statusList[i] == 0:
                statusList[i] = 1
                if desiredTotal - i <= 0 or not self.judge(desiredTotal - i, statusList, hmap):
                    statusList[i] = 0
                    hmap[key] = True
                    return True
                statusList[i] = 0
        hmap[key] = False
        return False

if __name__ == '__main__':
    solution = Solution()
    print(solution.canIWin(18, 79))
