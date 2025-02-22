from typing import Optional

from defs import ListNode, test


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(val=0, next=head)
        first, second = dummy, dummy

        for _ in range(n + 1):
            first = first.next

        while first:
            first = first.next
            second = second.next

        second.next = second.next.next

        return dummy.next

    # first calculating the length of the list and then traversing up to n - 1
    # redirect next pointer to skip the node we want to remove
    def removeNthFromEndLength(
        self, head: Optional[ListNode], n: int
    ) -> Optional[ListNode]:
        length = 0
        cur = head
        while cur:
            length += 1
            cur = cur.next

        if length == n:
            return head.next

        cur = head
        for _ in range(length - n - 1):
            cur = cur.next
        cur.next = cur.next.next

        return head

    TESTS = [
        test(
            ins=(ListNode.from_list([1, 2, 3, 4, 5]), 2),
            want=ListNode.from_list([1, 2, 3, 5]),
        ),
        test(
            ins=(ListNode.from_list([1]), 1),
            want=ListNode.from_list([]),
        ),
        test(
            ins=(ListNode.from_list([1, 2]), 1),
            want=ListNode.from_list([1]),
        ),
        test(
            ins=(ListNode.from_list([1, 2]), 2),
            want=ListNode.from_list([2]),
        ),
    ]
