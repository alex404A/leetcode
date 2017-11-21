class Solution(object):
    def __init__(self):
        self.statistics = {}
        self.pt = 0
        self.formula = ''

    def countOfAtoms(self, formula):
        """
        :type formula: str
        :rtype: str
        """
        if len(formula) == 0:
            return ''
        self.formula = formula
        nodes = self.parse()
        keys = sorted(nodes.iterkeys())
        result = ''
        for key in keys:
            number = nodes[key] if nodes[key] > 1 else ''
            result += (key + str(number))
        return result
    
    def parse(self):
        nodes = {}
        while self.pt < len(self.formula):
            if self.formula[self.pt] == ')':
                self.pt += 1
                cnt = self.parseNum()
                self.multiply(nodes, cnt)
                return nodes
            elif self.isUppercase(self.formula[self.pt]):
                word = self.parseLetter()
                cnt = self.parseNum()
                self.collectNode(nodes, word, cnt)
            elif self.formula[self.pt] == '(':
                self.pt += 1
                childNodes = self.parse()
                self.merge(nodes, childNodes)
        return nodes

    def parseNum(self):
        num = ''
        while self.pt < len(self.formula) and self.isNumber(self.formula[self.pt]):
            num += self.formula[self.pt]
            self.pt += 1
        return 1 if len(num) == 0 else int(num)

    def parseLetter(self):
        assert(self.isUppercase(self.formula[self.pt]) is True)
        word = self.formula[self.pt]
        self.pt += 1
        while self.pt < len(self.formula) and self.isLowercase(self.formula[self.pt]):
            word += self.formula[self.pt]
            self.pt += 1
        return word

    def isUppercase(self, a):
        diff = ord(a) - ord('A')
        return diff >= 0 and diff <= 25

    def isLowercase(self, a):
        diff = ord(a) - ord('a')
        return diff >= 0 and diff <= 25

    def isNumber(self, a):
        diff = ord(a) - ord('0')
        return diff >= 0 and diff <= 9

    def collectNode(self, nodes, word, cnt):
            if nodes.get(word) is None:
                nodes[word] = cnt
            else:
                nodes[word] += cnt
    
    def multiply(self, nodes, cnt):
        if cnt > 1:
            for key in nodes.iterkeys():
                nodes[key] = nodes[key] * cnt

    def merge(self, nodes, childNodes):
        for word, cnt in childNodes.iteritems():
            self.collectNode(nodes, word, cnt)

                
if __name__ == '__main__':
    solution = Solution()
    formula = "K4(ON(SO3)2)2"
    print(solution.countOfAtoms(formula))


