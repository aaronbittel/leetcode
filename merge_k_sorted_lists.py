from defs import ListNode, tests
from typing import Optional


class Solution:
    TESTS = [
        tests(ins=[[1, 4, 5], [1, 3, 4], [2, 6]], out=[1, 1, 2, 3, 4, 4, 5, 6]),
        tests(ins=[], out=[]),
        tests(ins=[[]], out=[]),
    ]

    def mergeKLists(self, lists: list[Optional[ListNode]]) -> Optional[ListNode]:
        return None
