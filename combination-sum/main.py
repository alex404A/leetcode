class Solution(object):
    def combinationSum(self, candidates, target):
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
        candidate = candidates[0]
        if target < candidate:
            return None
        if len(candidates) == 1:
            if target % candidate > 0:
                return None
            else:
                quotient = target / candidate
                return [[candidate] * quotient]
        else:
            result = []
            quotient = target / candidate
            for i in range(quotient + 1):
                tmp = self.getCombination(candidates[1:], target - candidate*i)
                if tmp is not None:
                    for j in range(len(tmp)):
                        tmp[j] += [candidate] * i
                    result += tmp
            return result

if __name__ == '__main__':
    solution = Solution()
    print(solution.combinationSum([2, 3, 1, 6, 7], 7))
    print(solution.combinationSum([2], 1))
