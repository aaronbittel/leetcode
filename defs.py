from __future__ import annotations

from collections import namedtuple
from typing import Optional

test = namedtuple("test", ["ins", "want"])


class ListNode:
    def __init__(self, val: int = 0, next: Optional[ListNode] = None) -> None:
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

    def __eq__(self, other: object, /) -> bool:
        if not isinstance(other, ListNode):
            return False

        cur = self
        cur_other = other
        while cur:
            if cur_other is None or cur.val != cur_other.val:
                return False

            cur = cur.next
            cur_other = cur_other.next

        return cur is None and cur_other is None

    def __str__(self) -> str:
        if not self:
            return "[]"

        out = f"{self.val}"

        cur = self
        while cur.next:
            cur = cur.next
            out += f"->{cur.val}"

        return out

    # TODO: do better
    def __repr__(self) -> str:
        return self.__str__()


class TreeNode:
    def __init__(
        self,
        val: int = 0,
        left: Optional[TreeNode] = None,
        right: Optional[TreeNode] = None,
    ) -> None:
        self.val = val
        self.left = left
        self.right = right

    def __repr__(self) -> str:
        if self is None:
            return "[]"

        as_list: list[Optional[int]] = []

        def traverse(head: Optional[TreeNode], acc: list[Optional[int]]):
            if head is None:
                return
            acc.append(head.val)
            if head.left is None and head.right is not None:
                acc.append(None)
            else:
                traverse(head.left, acc)
            if head.right is not None:
                traverse(head.right, acc)

        traverse(self, as_list)

        def join_optional_list(opt_list: list[Optional[int]]) -> str:
            out = "["
            for item in opt_list:
                if item is None:
                    out += "null, "
                else:
                    out += f"{item}, "
            return out[:-2] + "]"

        return join_optional_list(as_list)
