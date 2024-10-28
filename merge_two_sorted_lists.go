package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	empty_head := &ListNode{}
	cur := empty_head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur = addNode(cur, list1.Val)
			list1 = list1.Next
		} else {
			cur = addNode(cur, list2.Val)
			list2 = list2.Next
		}
	}

	if list1 == nil {
		for list2 != nil {
			cur = addNode(cur, list2.Val)
			list2 = list2.Next
		}
	}

	if list2 == nil {
		for list1 != nil {
			cur = addNode(cur, list1.Val)
			list1 = list1.Next
		}
	}

	return empty_head.Next
}

func addNode(cur *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value}
	cur.Next = newNode
	return newNode
}
