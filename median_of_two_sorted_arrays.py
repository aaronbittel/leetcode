from collections import namedtuple


class Solution:
    tests = namedtuple("Test", ["ins", "out"])
    TESTS = [
        tests(ins=([1, 3], [2]), out=2.0),
        tests(ins=([1, 2], [3, 4]), out=2.5),
        tests(ins=([1, 3, 4, 5, 7], [2, 3, 4, 6, 7, 9]), out=4.0),
        tests(ins=([1, 3, 5, 7], [2, 3, 4, 6, 7, 9]), out=4.5),
        tests(ins=([], [1]), out=1.0),
        tests(
            ins=([1, 2, 3, 4, 5], [6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17]), out=9.0
        ),
        tests(ins=([3], [-2, -1]), out=-1.0),
        tests(ins=([1, 2], [-1, 3]), out=1.5),
    ]

    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        n = len(nums1)
        m = len(nums2)

        l, r = 0, 0
        med, prev_med = 0, 0

        for _ in range((n + m) // 2 + 1):
            prev_med = med
            if n > l and m > r:
                if nums1[l] > nums2[r]:
                    med = nums2[r]
                    r += 1
                else:
                    med = nums1[l]
                    l += 1
            elif n <= l:
                med = nums2[r]
                r += 1
            elif m <= r:
                med = nums1[l]
                l += 1

        if (n + m) % 2 == 0:
            return (prev_med + med) / 2
        return med
