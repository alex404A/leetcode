class Solution(object):
    def __init__(self):
        results = []

    def multiply(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        self.results = [0 for i in range(len(num1) + len(num2))]
        for index1, d1 in enumerate(num1[::-1]):
            for index2, d2 in enumerate(num2[::-1]):
                d1Int = int(d1)
                d2Int = int(d2)
                self.updateResult(d1Int, d2Int, index1 + index2)

        cutIndex = 0
        for i in range(len(self.results) - 1, -1, -1):
            if self.results[i] != 0:
                cutIndex = i
                break

        results = self.results[0:cutIndex + 1]
        results.reverse()
        return ''.join([str(i) for i in results])

    def updateResult(self, d1, d2, index):
        result = d1 * d2 + self.results[index]
        self.results[index] = result % 10
        carry = result / 10
        index += 1
        while carry > 0:
            result = self.results[index] + carry
            self.results[index] = result % 10
            carry = result / 10
            index += 1

if __name__ == '__main__':
    solution = Solution()
    num1 = '231'
    num2 = '42'
    print(solution.multiply(num1, num2))
