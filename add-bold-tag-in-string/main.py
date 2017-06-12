class Solution(object):
    def __init__(self):
        self.startToWordDict = {}
        self.wordStartSet = set()

    def addBoldTag(self, rawStr, wordList):
        """
        :type s: str
        :type wordList: List[str]
        :rtype: str
        """
        if len(wordList) == 0:
            return s

        self.rawStr = rawStr
        self.wordList = wordList
        self.leastWordLen = self.wordList[0]
        for word in self.wordList[1:]:
            if len(word) < self.leastWordLen:
                self.leastWordLen = len(word)

        self.buildWordRelation()
        print(self.startToWordDict)
        print(self.wordStartSet)

    def buildWordRelation(self):

        def buildStartToWordDict():
            for word in self.wordList:
                wordStart = word[0:self.leastWordLen]
                if self.startToWordDict.get(wordStart) is None:
                    self.wordStartSet.add(wordStart)
                    self.startToWordDict[wordStart] = {}
                self.startToWordDict[wordStart][word] = []

        def addDepInDict():
            for word in self.wordList:
                wordStart = word[:self.leastWordLen]
                for i in range(len(word)):
                    depStartList = getDepStartList(word[i: i + self.leastWordLen])
                    self.startToWordDict[wordStart][word].extend(getDepWordList(word, i, depStartList))
                    self.startToWordDict[wordStart][word].sort(key = lambda dep: dep[0] + len(dep[2]), reverse=True)

        def getDepStartList(wordStart):
            result = []
            if len(wordStart) != self.leastWordLen:
                for start in self.wordStartSet:
                    if start.startswith(wordStart) is True:
                        result.append(start)
            elif wordStart in self.wordStartSet:
                result.append(wordStart)
            return result

        def getDepWordList(rawWord, index, depStartList):
            result = []
            for depStart in depStartList:
                if rawWord[0:self.leastWordLen] == depStart:
                    depWords = [(index, depStart, word) for word in self.startToWordDict[depStart].iterkeys() if len(word) > len(rawWord)]
                else:
                    depWords = [(index, depStart, word) for word in self.startToWordDict[depStart].iterkeys()]
                result.extend(depWords)
            return result

        buildStartToWordDict()
        addDepInDict()

    def getTagPosList(self):
        def getStartToWordListDict():
            result = {}
            for start in self.wordStartSet:
                result[start] = [word for word in self.startToWordDict[start].iterkeys()]
            return result

        def process(startOfWord, startIndex):
            def processOverlap(word, deps):
                for dep in deps:
                    partStr = self.rawStr[posDict['start'] + dep[0]:]
                    if dep[0] + len(dep[2]) > len(word) and partStr.startswith(dep[2]) is True:
                        posDict['end'] = posDict['start'] + len(dep[2]) + dep[0]
                        posDict['start'] = posDict['end']
                        processOverlap(dep[2], self.startToWordDict[dep[1]][dep[2]])
                        break

            posDict = {
                'start': startIndex,
                'end': startIndex
            }
            partStr = self.rawStr[posDict['start']:]
            for word in startToWords[startOfWord]:
                if partStr.startswith(word) is True:
                   posDict['end'] = posDict['start'] + len(word)
                   processOverlap(word, self.startToWordDict[startOfWord][word])
                   break
            return posDict['end']

        startToWords = getStartToWordListDict()
        tagPosList = []
        pt = 0
        while pt <= len(self.rawStr) - self.leastWordLen:
            start = self.rawStr[pt: pt + self.leastWordLen]
            wordDict = self.startToWordDict[start]
            if wordDict is None:
                pt += 1
                continue
            startPt = process(start, pt)
            if startPt == pt:
                pt += 1
            else:
                tagPosList.append((pt, startPt))
                pt = startPt
        return tagPosList

if __name__ == '__main__':
    solution = Solution()
    s = 'aaabbcc'
    wordList = ['aaa', 'aab', 'bc']
    solution.addBoldTag(s, wordList)
