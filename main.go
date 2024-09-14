package main

import "fmt"

func main() {
	head := &ListNode{Val: 3}
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 0}
	l3 := &ListNode{Val: -4}

	head.Next = l1
	l1.Next = l2
	l2.Next = l3
	l3.Next = l1

	fmt.Println(hasCycle(head))
}
