class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if len(strs) == 0:
            return ''
        else:
            return self.compressStrs(strs)

    def compressStrs(self, strs):
        result = []
        flag = False
        for i in range(0, len(strs), 2):
            if i == len(strs) - 1:
                common = strs[i]
            else:
                common = self.cmp2Str(strs[i], strs[i+1])
            if len(common) == 0:
                flag = True
                break
            else:
                result.append(common)
        print(result, flag)
        if not flag:
            if len(result) == 1:
                return result[0]
            else:
                return self.compressStrs(result)
        else:
            return ''

    def cmp2Str(self, str1, str2):
        length = min(len(str1), len(str2))
        commonPrefix = ''
        for i in range(0, length):
            if str1[i] == str2[i]:
                commonPrefix += str1[i]
            else:
                break
        return commonPrefix

if __name__ == '__main__':
    solution = Solution()
    test = ['hehe', 'hehehe', 'hehe', 'hehelalla', 'helala', 'hehouhou', 'he']
    test2 = ['']
    test3 = ['tianqingsedengyanyu']
    print(solution.longestCommonPrefix(test3))
