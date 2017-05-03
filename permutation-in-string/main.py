class Solution(object):
    def checkInclusion(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: bool
        """
        rawDict = self.getLetterDict(s1)
        copyDict = rawDict.copy()
        copyDict['letters'] = []
        for letter in s2:
            print(copyDict)
            num = copyDict.get(letter)
            if num is not None:
                if num > 0:
                    if copyDict.get('sum') == 1:
                        return True
                    copyDict[letter] -= 1
                    copyDict['sum'] -= 1
                    copyDict['letters'].append(letter)
                else:
                    self.genNewCopy(copyDict, letter)
            else:
                letters = copyDict.get('letters')
                if len(letters) > 0:
                    for letter in letters:
                        copyDict[letter] += 1
                        copyDict['sum'] += 1
                    copyDict['letters'] = []
        return False


    def getLetterDict(self, s):
        result = {'sum': 0}
        for letter in s:
            if result.get(letter) is None:
                result[letter] = 0
            result[letter] += 1
            result['sum'] += 1
        return result

    def genNewCopy(self, copyDict, letter):
        letters = copyDict.get('letters')
        index = letters.index(letter)
        for i in range(0, index):
            copyDict[letters[i]] += 1
            copyDict['sum'] += 1
        copyDict['letters'] = letters[index + 1:]
        copyDict['letters'].append(letter)

if __name__ == '__main__':
    solution = Solution()
    s1 = "adc"
    s2 = "adfghytdcda"
    print(solution.checkInclusion(s1, s2))
