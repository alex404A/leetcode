class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        strDict = {}
        for rawStr in strs:
            sortedStr = ''.join(sorted(rawStr))
            strList = strDict.get(sortedStr)
            if strDict.get(sortedStr) is None:
                strDict[sortedStr] = []
            strDict.get(sortedStr).append(rawStr)

        results = []
        for strList in strDict.itervalues():
            results.append(strList)
        return results

if __name__ == '__main__':
    solution = Solution()
    strs =["eat", "tea", "tan", "ate", "nat", "bat"]
    print(solution.groupAnagrams(strs))
