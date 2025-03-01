from defs import test


class Solution:
    def permute(self, nums: list[int]) -> list[list[int]]:
        res: list[list[int]] = []
        comb: list[int] = []

        def backtrack(numbers: list[int]) -> None:
            if not numbers:
                res.append(comb[:])
                return

            for i in range(len(numbers)):
                comb.append(numbers[i])
                backtrack(numbers[:i] + numbers[i + 1 :])
                comb.pop()

        backtrack(nums)
        return res

    TESTS = [
        test(
            ins=[1, 2, 3],
            want=[[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]],
        ),
        test(ins=[0, 1], want=[[0, 1], [1, 0]]),
        test(ins=[1], want=[[1]]),
    ]
