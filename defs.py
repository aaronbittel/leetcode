from __future__ import annotations

from collections import namedtuple
from typing import Optional

tests = namedtuple("tests", ["ins", "out"])


class ListNode:
    def __init__(self, val: int = 0, next: Optional[ListNode] = None):
        self.val = val
        self.next = next
