class Solution(object):
    def __init__(self):
        self.statistics = {}
        self.pt = 0

    def countOfAtoms(self, formula):
        """
        :type formula: str
        :rtype: str
        """
        if (len(formula) == 0):
            return ''
        return self.parse(formula)
    
    def parse(self, formula):
        while (self.pt < len(formula))
        if (self.isUppercase(formula[self.pt])):

    def parseBracketBlk(self, formula):
        pass

    def isUppercase(self, a):
        diff = ord(a) - ord('A')
        return diff >= 0 and diff <= 25

    def isLowercase(self, a):
        diff = ord(a) - ord('a')
        return diff >= 0 and diff <= 25

    def isNumber(self, a):
        diff = ord(a) - ord('0')
        return diff >= 0 and diff <= 9

