package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(items []int) *ListNode {
	switch len(items) {
	case 0:
		return nil
	case 1:
		return &ListNode{
			Val:  items[0],
			Next: nil,
		}
	}

	head := &ListNode{
		Val:  items[0],
		Next: nil,
	}

	cur := head

	for i := 1; i < len(items); i++ {
		cur.Next = &ListNode{
			Val:  items[i],
			Next: nil,
		}
		cur = cur.Next
	}

	return head
}

func (l *ListNode) String() string {
	if l == nil {
		return "{ }"
	}
	out := ""
	cur := l

	for cur != nil {
		out += fmt.Sprintf("%d -> ", cur.Val)
		cur = cur.Next
	}

	return out
}
