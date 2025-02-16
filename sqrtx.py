from defs import test


class Solution:
    TESTS = [
        test(ins=4, want=2),
        test(ins=5, want=2),
        test(ins=8, want=2),
        test(ins=9, want=3),
        test(ins=128, want=11),
        test(ins=1, want=1),
        test(ins=2, want=1),
        test(ins=36, want=6),
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
