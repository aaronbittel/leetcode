from defs import ListNode


def test_listnode_equal_true():
    tests = [
        (None, None),
        (ListNode(1), ListNode(1)),
        (ListNode.from_list([1, 2, 3]), ListNode.from_list([1, 2, 3])),
    ]

    for tt in tests:
        assert tt[0] == tt[1]


def test_listnode_equal_false():
    tests = [
        (ListNode(2), ListNode(1)),
        (None, ListNode(1)),
        (ListNode.from_list([1, 2, 3]), ListNode.from_list([1, 2])),
    ]

    for tt in tests:
        assert tt[0] != tt[1]
