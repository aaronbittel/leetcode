from defs import test


class Solution:
    def merge(self, intervals: list[list[int]]) -> list[list[int]]:
        if not intervals:
            return []

        intervals.sort(key=lambda i: i[0])

        res: list[list[int]] = [intervals[0]]

        for i in range(1, len(intervals)):
            if res[-1][1] < intervals[i][0]:
                res.append(intervals[i])
            else:
                res[-1][0] = min(res[-1][0], intervals[i][0])
                res[-1][1] = max(res[-1][1], intervals[i][1])

        return res

    TESTS = [
        test(ins=[[1, 3], [2, 6], [8, 10], [15, 18]], want=[[1, 6], [8, 10], [15, 18]]),
        test(ins=[[1, 4], [4, 5]], want=[[1, 5]]),
        test(ins=[[1, 4], [0, 4]], want=[[0, 4]]),
        test(ins=[[1, 4], [0, 1]], want=[[0, 4]]),
        test(ins=[[1, 4], [0, 0]], want=[[0, 0], [1, 4]]),
    ]


Solution().merge(Solution.TESTS[0].ins)
