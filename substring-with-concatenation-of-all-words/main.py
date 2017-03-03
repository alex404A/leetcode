class Solution(object):
    def __init__(self):
        self.results = []
        self.wordLen = 0
        self.strLen = 0

    def findSubstring(self, s, words):
        """
        :type s: str
        :type words: List[str]
        :rtype: List[int]
        """

        self.strLen = len(s)
        if len(words) == 0:
            return [i for i in range(0, self.strLen)]

        self.wordLen = len(words[0])
        if self.wordLen > self.strLen:
            return []

        wordPositionDict = self.findPositionOfEveryWord(s, words)
        print(wordPositionDict)
        if self.isAnyWordNotEnough(words, wordPositionDict):
            return []

        wordPositionsTupleList = sorted(wordPositionDict.items(), key = lambda item: len(item[1]))
        minimumAppearanceWord = wordPositionsTupleList[0]
        remainingWords = [w for i, w in enumerate(words) if i != words.index(minimumAppearanceWord[0])]
        print(minimumAppearanceWord, remainingWords, wordPositionDict)
        for pos in minimumAppearanceWord[1]:
            self.findCompleteSubstring(pos, pos+self.wordLen, s, remainingWords, wordPositionDict)
        print(self.results)
        return list(set(self.results))

    def findPositionOfEveryWord(self, s, words):
        wordPositionDict = dict((word, []) for word in words)
        for i in range(len(s) - self.wordLen + 1):
            wordPositionList = wordPositionDict.get(s[i:i+self.wordLen])
            if wordPositionList is not None:
                wordPositionList.append(i)
        return wordPositionDict

    def genWordNumDict(self, words):
        wordNumDict = {}
        for word in words:
            if wordNumDict.get(word) is None:
                wordNumDict[word] = 0
            wordNumDict[word] += 1
        return wordNumDict

    def isAnyWordNotEnough(self, words, wordPositionDict):
        wordNumDict = self.genWordNumDict(words)
        for k, v in wordPositionDict.iteritems():
            if len(v) < wordNumDict.get(k):
                return True
        return False

    # left refers to the position of left boundary
    # right refers to the position of right boundary plus 1
    def findCompleteSubstring(self, left, right, s, remainingWords, wordPositionDict):
        if len(remainingWords) == 0:
            self.results.append(left)
            return
        if left >= self.wordLen:
            word = s[left-self.wordLen: left]
            if self.isAnyWordMatchedInSpecificPos(left-self.wordLen, word, remainingWords, wordPositionDict):
                newRemainingWords = [w for i, w in enumerate(remainingWords) if i != remainingWords.index(word)]
                self.findCompleteSubstring(left-self.wordLen, right, s, newRemainingWords, wordPositionDict)
        if right < len(s) - self.wordLen + 1:
            word = s[right: right+self.wordLen]
            if self.isAnyWordMatchedInSpecificPos(right, word, remainingWords, wordPositionDict):
                newRemainingWords = [w for i, w in enumerate(remainingWords) if i != remainingWords.index(word)]
                self.findCompleteSubstring(left, right+self.wordLen, s, newRemainingWords, wordPositionDict)

    def isAnyWordMatchedInSpecificPos(self, position, word, remainingWords, wordPositionDict):
        if word in remainingWords and position in wordPositionDict.get(word):
            return True
        return False


if __name__ == '__main__':
    solution = Solution()
    s = 'wordgoodgoodgoodbestword'
    words = ["word","good","best","good"]
    solution.findSubstring(s, words)
