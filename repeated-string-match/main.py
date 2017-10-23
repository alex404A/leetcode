class Solution(object):
    def repeatedStringMatch(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: int
        """
        aLen = len(A)
        cnt = 0
        if B.find(A) > -1:
            start = B.find(A)
            end = len(B) - 1
            pt = start
            while pt < len(B):
                sliceItem = B[pt:pt + aLen]
                if sliceItem == A:
                    cnt += 1
                    end = pt + aLen
                pt += aLen
            B = B[0:start] + B[end:]
        if len(B) == 0:
            return cnt
        pt = 0
        while pt < len(B):
            start = A.find(B[0:pt])
            end = A.find(B[pt:])
            if start > -1 and end > -1 and self.isContinius(start, end, pt, aLen):
                return cnt + 2 if pt != 0 else cnt + 1
            else:
                pt += 1
        return -1

    def isContinius(self, start, end, pt, length):
        if start + pt >= length:
            return start + pt - length == end
        else:
            return start + pt == end

if __name__ == '__main__':
    solution = Solution()
    A = "abcd"
    B = "cdabcdab"
    print(solution.repeatedStringMatch(A, B))
                
            
        
