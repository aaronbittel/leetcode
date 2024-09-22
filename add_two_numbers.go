package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carryOver := 0

	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carryOver
		cur.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		carryOver = sum / 10

		cur = cur.Next
	}

	if carryOver > 0 {
		cur.Next = &ListNode{
			Val:  carryOver,
			Next: nil,
		}
	}

	return dummy.Next
}

func addTwoNumbersExtraNumArray(l1 *ListNode, l2 *ListNode) *ListNode {
	sumArr := []int{}
	carryOver := 0

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carryOver
		sumArr = append(sumArr, sum%10)
		carryOver = sum / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		for l2 != nil {
			sum := (l2.Val + carryOver) % 10
			sumArr = append(sumArr, sum)
			carryOver = (l2.Val + carryOver) / 10
			l2 = l2.Next
		}
	} else {
		for l1 != nil {
			sum := (l1.Val + carryOver) % 10
			sumArr = append(sumArr, sum)
			carryOver = (l1.Val + carryOver) / 10
			l1 = l1.Next
		}
	}

	if carryOver > 0 {
		sumArr = append(sumArr, carryOver)
	}

	return NewLinkedList(sumArr)
}

func addTwoNumbersNumbersTooBig(l1 *ListNode, l2 *ListNode) *ListNode {
	getNumber := func(l *ListNode) int {
		number := 0
		cur := l
		for i := 1; cur != nil; i *= 10 {
			number += i * cur.Val
			cur = cur.Next
		}
		return number
	}

	fmt.Println(getNumber(l1), getNumber(l2))
	sum := getNumber(l1) + getNumber(l2)
	fmt.Println(sum)

	if sum == 0 {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	sumArr := []int{}
	for sum != 0 {
		sumArr = append(sumArr, sum%10)
		sum /= 10
	}

	return NewLinkedList(sumArr)
}
