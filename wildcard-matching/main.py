class Solution(object):
    def __init__(self):
        self.status = {}

    def isMatch(self, s, p):
        p1 = [i for i in p if i != '*']
        if len(p1) > len(s):
            return False
        return self.isOK(s, p)

    def isOK(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        if p == '' and s == '':
            return True
        if p == '' and len(s) > 0:
            return False
        if p[0] == '?':
            if len(s) > 0 and self.checkStatus(s[1:], p[1:]):
                return True
        elif p[0] == '*':
            if self.checkStatus(s, p[1:]):
                return True
            if len(s) > 0 and self.checkStatus(s[1:], p):
                return True
            if len(s) > 0 and self.checkStatus(s[1:], p[1:]):
                return True
        else:
            if len(s) > 0 and s[0] == p[0] and self.checkStatus(s[1:], p[1:]):
                return True
        self.status[(len(s), len(p))] = False
        return False

    def checkStatus(self, s, p):
        status = self.status.get((len(s), len(p)))
        if status is not None:
            return status
        status = self.isOK(s, p)
        self.status[(len(s), len(p))] = status
        return status


if __name__ == '__main__':
    solution = Solution()
    s = ''.join(['a'] * 1000)
    p = '*' + ''.join(['a'] * 1002)
    print(s, p)
    print(solution.isMatch(s, p))
