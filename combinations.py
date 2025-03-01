from defs import test


class Solution:
    def combine(self, n: int, k: int) -> list[list[int]]:
        res = []
        comb = []

        def backtrack(start: int) -> None:
            if len(comb) == k:
                res.append(comb[:])
                return

            for num in range(start, n + 1):
                comb.append(num)
                backtrack(num + 1)
                comb.pop()

        backtrack(1)
        return res

    TESTS = [
        test(ins=(4, 2), want=[[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]),
        test(ins=(1, 1), want=[[1]]),
    ]
