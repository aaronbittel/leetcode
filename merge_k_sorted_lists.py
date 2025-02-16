from defs import ListNode, test
from typing import Optional


class Solution:
    TESTS = [
        test(
            ins=[
                ListNode.from_list([1, 4, 5]),
                ListNode.from_list([1, 3, 4]),
                ListNode.from_list([2, 6]),
            ],
            want=ListNode.from_list([1, 1, 2, 3, 4, 4, 5, 6]),
        ),
        test(ins=None, want=None),
        test(ins=[None], want=None),
    ]

    def mergeKLists(self, lists: list[Optional[ListNode]]) -> Optional[ListNode]:
        if not lists or len(list(l for l in lists if l)) == 0:
            print("returned early")
            return None

        head = None
        cur = head

        while len(list(l for l in lists if l)) > 0:
            min_val: Optional[int] = None
            min_idx: Optional[int] = None
            for i, l in enumerate(lists):
                # print(f"{i = }")
                if l is None:
                    continue

                if min_val is None:
                    min_val = l.val
                    min_idx = i
                    # print(f"first set: {min_val = }, {min_idx = }")
                    continue

                if l.val < min_val:
                    min_val = l.val
                    min_idx = i
                    # print(f"found smaller: {min_val = }, {min_idx = }")

            if min_val is not None and min_idx is not None:
                if head is None:
                    head = ListNode(val=min_val)
                    cur = head
                else:
                    ln = ListNode(val=min_val)
                    cur.next = ln
                    cur = ln
                if lists[min_idx] is not None:
                    lists[min_idx] = lists[min_idx].next
                print(f"{min_val = }, {min_idx = }")
                print(f"cur head: {head.__str__() = }")

        return head


if __name__ == "__main__":
    for tt in Solution.TESTS:
        got = Solution().mergeKLists(tt.ins)
        if got != tt.want:
            out = ""
            for l in tt.ins:
                out += f"\t{l}\n"
            print(f"Ins: {out}FAILURE! Expected: {tt.want}, got: {got}\n")
        else:
            print(f"Ins: {tt.ins} -> SUCCESS! {got}")
