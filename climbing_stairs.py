from functools import cache

from defs import test


class Solution:
    TESTS = [
        test(ins=2, want=2),
        test(ins=3, want=3),
        test(ins=38, want=63245986),
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
