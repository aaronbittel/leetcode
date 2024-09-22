package main

import (
	"fmt"
)

func main() {
	l1 := NewLinkedList([]int{2, 4, 3})
	l2 := NewLinkedList([]int{5, 6, 4})
	// fmt.Println(addTwoNumbers(l1, l2))
	//
	// l1 = NewLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	// l2 = NewLinkedList([]int{5, 6, 4})
	// fmt.Println(addTwoNumbers(l1, l2))

	l1 = NewLinkedList([]int{2, 4, 9})
	l2 = NewLinkedList([]int{5, 6, 4, 9})
	fmt.Println(addTwoNumbers(l1, l2))

	// l1 = NewLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	// l2 = NewLinkedList([]int{9, 9, 9, 9})
	// fmt.Println(addTwoNumbers(l1, l2))
	//
	// l1 = NewLinkedList([]int{3, 7})
	// l2 = NewLinkedList([]int{9, 2})
	// fmt.Println(addTwoNumbers(l1, l2))
}
