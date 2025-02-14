from functools import cache


class Solution:
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


from collections import namedtuple  # noqa: E402

tests = namedtuple("Test", ["ins", "out"])
tests = [
    tests(ins=2, out=2),
    tests(ins=3, out=3),
    tests(ins=38, out=3),
]

for tt in tests:
    got = Solution().climbStairs(tt.ins)
    if got != tt.out:
        print(f"Ins: {tt.ins} -> Expected: {tt.out}, got: {got}")
    else:
        print(f"Ins: {tt.ins} -> {got}: Success!")
