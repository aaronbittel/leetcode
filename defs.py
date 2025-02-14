from __future__ import annotations

from collections import namedtuple
from typing import Optional

tests = namedtuple("tests", ["ins", "out"])


class ListNode:
    def __init__(self, val: int = 0, next: Optional[ListNode] = None):
        self.val = val
        self.next = next

    @staticmethod
    def from_list(l: list[int]) -> Optional[ListNode]:
        if not l:
            return None

        head = ListNode(val=l[0], next=None)
        cur = head
        for v in l[1:]:
            ln = ListNode(val=v, next=None)
            cur.next = ln
            cur = ln

        return head

    def __str__(self) -> str:
        if not self:
            return "[]"

        out = f"{self.val}"

        cur = self
        while cur.next:
            cur = cur.next
            out += f"->{cur.val}"

        return out
