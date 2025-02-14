from collections import namedtuple


class Solution:
    tests = namedtuple("Test", ["ins", "out"])
    TESTS = [
        tests(ins=4, out=2),
        tests(ins=5, out=2),
        tests(ins=8, out=2),
        tests(ins=9, out=3),
        tests(ins=128, out=11),
        tests(ins=1, out=1),
        tests(ins=2, out=1),
        tests(ins=36, out=6),
    ]

    def mySqrt(self, x: int) -> int:
        lo, hi = 0, x
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            p = mid * mid
            if p == x:
                return mid

            if p > x:
                hi = mid - 1
            else:
                lo = mid + 1

        return lo - 1
