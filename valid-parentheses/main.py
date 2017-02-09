class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        parenthesesList = []
        leftParentheses = ['{', '[', '(']
        rightParentheses = ['}', ']', ')']
        result = True
        for i in range(len(s)):
            if s[i] in leftParentheses:
                parenthesesList.append(s[i])
            if s[i] in rightParentheses:
                if len(parenthesesList) == 0:
                    result = False
                    break
                else:
                    left = parenthesesList[len(parenthesesList)-1]
                    right = s[i]
                    if leftParentheses.index(left) == \
                            rightParentheses.index(right):
                        parenthesesList = parenthesesList[0: -1]
                    else:
                        result = False
                        break
        if len(parenthesesList) > 0:
            return False
        return result
