class Solution:
    # @return an integer
    def divide(self, dividend, divisor):
        limit = 2147483647
        isNegative = (dividend < 0 and divisor > 0) or (dividend > 0 and divisor < 0)
        if isNegative:
            if abs(dividend) < abs(divisor):
                return 0
        sum = 0; count = 0; res = 0
        a = abs(dividend); b = abs(divisor)
        while a >= b:
            sum = b
            count = 1
            while sum + sum <= a:
                sum += sum
                count += count
            a -= sum
            res += count
            if res > limit:
                return 0 - limit if isNegative else limit
        if isNegative:
            res = 0 - res if a == 0 else 0 - res - 1
        return res


if __name__ == '__main__':
    solution = Solution()
    print(solution.divide(1026117192, -874002063))
