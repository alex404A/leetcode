import time

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
            return rawStr

        self.rawStr = rawStr
        self.wordList = wordList
        self.leastWordLen = len(self.wordList[0])
        for word in self.wordList[1:]:
            if len(word) < self.leastWordLen:
                self.leastWordLen = len(word)

        self.buildWordRelation()
        # print(self.startToWordDict)
        # print(self.wordStartSet)
        tagPosList = self.getTagPosList()
        return self.addTags(tagPosList)

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
                    # self.startToWordDict[wordStart][word].sort(key = lambda dep: dep[0] + len(dep[2]), reverse=True)

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
                if rawWord[0:self.leastWordLen] == depStart and index == 0:
                    depWords = [(index, depStart, word) for word in self.startToWordDict[depStart].iterkeys() if word != rawWord and index + len(word) > len(rawWord)]
                else:
                    depWords = [(index, depStart, word) for word in self.startToWordDict[depStart].iterkeys() if index + len(word) > len(rawWord)]
                result.extend(depWords)
            return result

        buildStartToWordDict()
        addDepInDict()

    def getTagPosList(self):
        def getStartToWordListDict():
            result = {}
            for start in self.wordStartSet:
                result[start] = [word for word in self.startToWordDict[start].iterkeys()]
                result[start].sort(key = lambda word: len(word), reverse=True)
            return result

        def process(startOfWord, startIndex):
            def processOverlap(word, deps):
                for dep in deps:
                    partStr = self.rawStr[posDict['start'] + dep[0]:]
                    if dep[0] + len(dep[2]) > len(word) and partStr.startswith(dep[2]) is True:
                        tmp = posDict['start'] + dep[0]
                        posDict['end'] = tmp + len(dep[2])
                        posDict['start'] = tmp
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
            wordDict = self.startToWordDict.get(start)
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

    def addTags(self, tagPosList):
        newTagPosList = []

        def mergeAdjacentTag():
            if len(tagPosList) == 0:
                return
            curStart = tagPosList[0][0]
            for i in range(len(tagPosList)):
                if i + 1 < len(tagPosList) and tagPosList[i][1] != tagPosList[i + 1][0]:
                    newTagPosList.append((curStart, tagPosList[i][1]))
                    curStart = tagPosList[i + 1][0]
                elif i + 1 == len(tagPosList):
                    newTagPosList.append((curStart, tagPosList[i][1]))

        mergeAdjacentTag()
        formattedStr = self.rawStr
        offset = 0
        for start, end in newTagPosList:
            formattedStr = formattedStr[:start + offset] + '<b>' + formattedStr[start + offset:]
            offset += 3
            formattedStr = formattedStr[:end + offset] + '</b>' + formattedStr[end + offset:]
            offset += 4
        return formattedStr

if __name__ == '__main__':
    solution = Solution()
