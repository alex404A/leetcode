import re

class Solution(object):
    def __init__(self):
        self.startTagPattern = re.compile(r"<([A-Z]{1,9})>")
        self.endTagPattern = re.compile(r"</([A-Z]{1,9})>")
        self.startCdataPattern = re.compile(r"<\!\[CDATA\[")
        self.endCdataPattern = re.compile(r"\]\]>")
        self.maxTagLen = 12
        self.tokens = []

    def isValid(self, code):
        """
        :type code: str
        :rtype: bool
        """
        code = code.strip()
        if self.startTagPattern.match(code) is None:
            return False
        i = 0
        while i < len(code):
            if code[i] == '<':
                startTagMatch = self.startTagPattern.match(code, i)
                endTagMatch = self.endTagPattern.match(code, i)
                startCdataMatch = self.startCdataMatch.match(code, i)
            if startTagMatch is not None:
                i = self.processStartTagMatch(startTagMatch)
                continue
            if endTagMatch is not None:
                i, isValid = self.processEndTagMatch(endTagMatch)
                if isValid:
                    continue
                return False
            if startCdataMatch is not None:
                i = self.processStartCdataMatch(startCdataMatch)
                endCdataMatch = self.endCdataPattern.match(code, i)
                if endCdataPattern is None:
                    return False
                i = self.processEndDataMatch(endTagPattern)
                continue
            i += 1
        return True

    def processStartTagMatch(self, match):
        pass

    def processEndTagMatch(self, match):
        pass

    def processStartCdataMatch(self, match):
        pass

    def processEndDataMatch(self, match):
        pass
