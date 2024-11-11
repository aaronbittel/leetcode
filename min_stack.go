package main

import (
	"math"
)

// LinkedList with extra min
// Stores the curMinValue at the point when the value was pushed onto the stack
// No need for extra storing a sorted list of the elements
// No need for extra calculating the minVal when the minVal is poped

type Elem struct {
	val  int
	min  int
	next *Elem
}

type MinStack struct {
	head *Elem
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &Elem{val: val, min: val}
	} else {
		this.head = &Elem{val: val, min: min(val, this.head.min), next: this.head}
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

// Pros:
// 	- Only need to store the elements once + minVal
// 	- Pop: If last element is not the minVal, then Pop does not require any minVal computation
// Cons:
// 	- Pop: If last element is the minVal, need to traverse the all elements to find the new minVal

// Alternative: Keep a sorted list of the elements -> double the size
type MinStack1 struct {
	items  []int
	minVal int
}

func NewMinStack1() MinStack1 {
	return MinStack1{
		items:  make([]int, 0),
		minVal: MaxInt(),
	}
}

func (this *MinStack1) Push1(val int) {
	if val < this.minVal {
		this.minVal = val
	}
	this.items = append(this.items, val)
}

func (this *MinStack1) Pop1() {
	last := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	if last == this.minVal {
		this.minVal = MaxInt()
		for i := 0; i < len(this.items); i++ {
			if this.items[i] < this.minVal {
				this.minVal = this.items[i]
			}
		}
	}
}

func (this *MinStack1) Top1() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack1) GetMin1() int {
	return this.minVal
}

func MaxInt() int {
	return int(math.MaxInt >> 1)
}
