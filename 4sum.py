from defs import test


class Solution:
    def fourSum(self, nums: list[int], target: int) -> list[list[int]]:
        nums.sort()

        n = len(nums)
        res: list[list[int]] = []

        for L in range(n):
            if L > 0 and nums[L] == nums[L - 1]:
                continue
            for R in range(L + 1, n):
                if R > L + 1 and nums[R] == nums[R - 1]:
                    continue
                l, r = R + 1, n - 1
                while l < r:
                    sum = nums[L] + nums[R] + nums[l] + nums[r]
                    if sum == target:
                        res.append([nums[L], nums[R], nums[l], nums[r]])
                        l += 1
                        r -= 1
                        while l < r and nums[l] == nums[l - 1]:
                            l += 1
                        while l < r and nums[r] == nums[r + 1]:
                            r -= 1
                    elif sum > target:
                        r -= 1
                    else:
                        l += 1

        return res

    TESTS = [
        test(
            ins=([1, 0, -1, 0, -2, 2], 0),
            want=[[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]],
        ),
        test(ins=([2, 2, 2, 2, 2], 8), want=[[2, 2, 2, 2]]),
        test(ins=([-3, -1, 0, 2, 4, 5], 0), want=[[-3, -1, 0, 4]]),
        test(
            ins=([-3, -2, -1, 0, 0, 1, 2, 3], 0),
            want=[
                [-3, -2, 2, 3],
                [-3, -1, 1, 3],
                [-3, 0, 0, 3],
                [-3, 0, 1, 2],
                [-2, -1, 0, 3],
                [-2, -1, 1, 2],
                [-2, 0, 0, 2],
                [-1, 0, 0, 1],
            ],
        ),
    ]
