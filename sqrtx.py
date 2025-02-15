from defs import tests


class Solution:
    TESTS = [
        tests(ins=4, want=2),
        tests(ins=5, want=2),
        tests(ins=8, want=2),
        tests(ins=9, want=3),
        tests(ins=128, want=11),
        tests(ins=1, want=1),
        tests(ins=2, want=1),
        tests(ins=36, want=6),
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
