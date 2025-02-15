from typing import Optional

from defs import ListNode, tests


class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if head is None:
            return None

        if k == 1:
            return head

        batches, leftover = batch(head, k)

        for i in range(len(batches)):
            batches[i] = reverse(batches[i], k - 1)

        result = batches[0]

        for b in batches[1:]:
            last(result).next = b

        if leftover is not None:
            last(result).next = leftover

        return result

    TESTS = [
        tests(
            ins=(ListNode.from_list([1, 2, 3, 4, 5]), 2),
            want=ListNode.from_list([2, 1, 4, 3, 5]),
        ),
        tests(
            ins=(ListNode.from_list([1, 2, 3, 4]), 2),
            want=ListNode.from_list([2, 1, 4, 3]),
        ),
        tests(
            ins=(ListNode.from_list([1, 2, 3, 4, 5]), 3),
            want=ListNode.from_list([3, 2, 1, 4, 5]),
        ),
        tests(
            ins=(ListNode.from_list([1, 2]), 2),
            want=ListNode.from_list([2, 1]),
        ),
    ]


def split(
    head: Optional[ListNode], length: int
) -> tuple[Optional[ListNode], Optional[ListNode]]:
    if head is None:
        return None, None
    cur = head
    for _ in range(length - 1):
        if cur.next is None:
            return None, head
        cur = cur.next
    rest = cur.next
    cur.next = None
    return head, rest


def batch(head: ListNode, length: int) -> tuple[list[ListNode], Optional[ListNode]]:
    batches: list[ListNode] = []
    while True:
        part, rest = split(head, length)
        match (part, rest):
            case (None, None):
                return batches, None
            case (part, None):
                batches.append(part)
                return batches, None
            case (None, left):
                return batches, left
            case (part, rest):
                batches.append(part)
                head = rest


def last_to_first(head: ListNode) -> ListNode:
    prev_last = head

    while prev_last.next.next:
        prev_last = prev_last.next

    last = prev_last.next
    prev_last.next = None
    last.next = head
    return last


def last(head: ListNode) -> ListNode:
    cur = head
    while cur.next:
        cur = cur.next
    return cur


def reverse(head: ListNode, length: int) -> ListNode:
    if length == 0:
        return head

    head = last_to_first(head)

    cur = head.next
    head.next = reverse(cur, length - 1)
    return head
