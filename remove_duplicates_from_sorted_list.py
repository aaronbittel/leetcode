from defs import ListNode, test
from typing import Optional


class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return None

        h = ListNode(val=head.val)
        cur_h = h

        cur = head
        last = cur.val
        while cur:
            if cur.val != last:
                cur_h.next = ListNode(val=cur.val)
                cur_h = cur_h.next
                last = cur.val
            cur = cur.next

        return h

    TESTS = [
        test(ins=ListNode.from_list([1, 1, 2]), want=ListNode.from_list([1, 2])),
        test(
            ins=ListNode.from_list([1, 1, 2, 3, 3]), want=ListNode.from_list([1, 2, 3])
        ),
    ]
