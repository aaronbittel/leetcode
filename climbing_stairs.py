from collections import namedtuple
from functools import cache


class Solution:
    tests = namedtuple("Test", ["ins", "out"])
    TESTS = [
        tests(ins=2, out=2),
        tests(ins=3, out=3),
        tests(ins=38, out=63245986),
    ]

    @cache
    def climbStairs(self, n: int) -> int:
        if n == 0 or n == 1:
            return 1

        if n < 0:
            return 0

        total = 0

        total += self.climbStairs(n - 1)
        total += self.climbStairs(n - 2)

        return total
