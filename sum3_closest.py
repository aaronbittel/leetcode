# Solution
# 3sum!!!
# 1. sort numbers
# 2. go through each number
#   - l (first after cur), r (last) pointer
#   - cal cur_sum, cal diff
#   - if diff < 0 -> l++ (make sum bigger) else r-- (make sum smaller)
# Runtime: O(n**2) -> n * (n - i)

from defs import tests


class Solution:
    TESTS = [
        tests(ins=([-1, 2, 1, -4], 1), want=2),
        tests(ins=([0, 0, 0], 1), want=0),
        tests(ins=([0, 1, 2], 0), want=3),
    ]

    def threeSumClosest(self, nums: list[int], target: int) -> int:
        nums.sort()

        min_sum = sum(nums[:3])
        min_diff = abs(target - min_sum)

        for i, n in enumerate(nums):
            l, r = i + 1, len(nums) - 1
            while l < r:
                cur_sum = n + nums[l] + nums[r]
                diff = cur_sum - target

                if diff < 0:
                    l += 1
                    diff = -diff
                else:
                    r -= 1

                if diff < min_diff:
                    min_diff = diff
                    min_sum = cur_sum

        return min_sum
