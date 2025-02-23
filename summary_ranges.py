from defs import test


class Solution:
    def summaryRanges(self, nums: list[int]) -> list[str]:
        if not nums:
            return []

        res: list[str] = []

        start_i = 0
        for i in range(len(nums)):
            if i == len(nums) - 1 or nums[i] + 1 != nums[i + 1]:
                if start_i == i:
                    res.append(f"{nums[i]}")
                else:
                    res.append(f"{nums[start_i]}->{nums[i]}")
                start_i = i + 1

        return res

    TESTS = [
        test(ins=[0, 1, 2, 4, 5, 7], want=["0->2", "4->5", "7"]),
        test(ins=[0, 2, 3, 4, 6, 8, 9], want=["0", "2->4", "6", "8->9"]),
    ]
