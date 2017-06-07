import re

class Solution(object):
    def __init__(self):
        self.startTagPattern = re.compile(r"<([A-Z]{1,9})>")
        self.endTagPattern = re.compile(r"</([A-Z]{1,9})>")
        self.startCdataPattern = re.compile(r"<\!\[CDATA\[")
        self.endCdataPattern = re.compile(r"\]\]>")
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
                startCdataMatch = self.startCdataPattern.match(code, i)
                if startTagMatch is not None:
                    i = self.processStartTagMatch(startTagMatch)
                    continue
                if endTagMatch is not None:
                    i = self.processEndTagMatch(endTagMatch)
                    if i == -1:
                        return False
                    if len(self.tokens) == 0 and i < len(code):
                        return False
                    continue
                if startCdataMatch is not None:
                    i = self.processStartCdataMatch(startCdataMatch)
                    endCdataMatch = self.endCdataPattern.search(code, i)
                    if endCdataMatch is None:
                        return False
                    i = self.processEndDataMatch(endCdataMatch)
                    continue
                return False
            i += 1
        return True if len(self.tokens) == 0 else False

    def processStartTagMatch(self, match):
        self.tokens.append(match.group(1))
        return match.end()

    def processEndTagMatch(self, match):
        endTagName = match.group(1)
        if len(self.tokens) > 0 and endTagName == self.tokens.pop():
            return match.end()
        return -1

    def processStartCdataMatch(self, match):
        return match.end()

    def processEndDataMatch(self, match):
        return match.end()

if __name__ == '__main__':
    solution = Solution()
    code = '<A>  </A><B></B>'
    print(solution.isValid(code))
