class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        candidates = sorted([c for c in candidates if c <= target])
        if len(candidates) == 0:
            return []
        result = self.getCombination(candidates, target)
        return [] if result is None else result

    def getCombination(self, candidates, target):
        if target == 0:
            return [[]]
        if len(candidates) == 0:
            return None
        candidate = candidates[0]
        if target < candidate:
            return None
        if len(candidates) == 1:
            if target != candidate:
                return None
            else:
                return [[candidate]]
        else:
            result = []
            cnt = self.getRepeatCount(candidates)
            for i in range(cnt+1):
                tmp = self.getCombination(candidates[cnt:], target - candidate*i)
                if tmp is not None:
                    for j in range(len(tmp)):
                        tmp[j] += [candidate] * i
                    result += tmp
            return result

    def getRepeatCount(self, candidates):
        cnt = 1
        for i in range(1, len(candidates)):
            if candidates[i] == candidates[i-1]:
                cnt += 1
            else:
                break;
        return cnt



if __name__ == '__main__':
    solution = Solution()
    print(solution.combinationSum2([1, 1], 1))
