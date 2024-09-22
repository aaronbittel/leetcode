package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	set := map[*ListNode]struct{}{head: {}}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
		if _, found := set[cur]; found {
			return true
		}
		set[cur] = struct{}{}
	}
	return false
}
