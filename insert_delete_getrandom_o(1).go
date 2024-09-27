package main

import (
	"math/rand"
	"slices"
)

// faster, but more memory heavy (154ms, 50.78mb)
type RandomizedSet struct {
	set    map[int]struct{}
	values []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set:    make(map[int]struct{}),
		values: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if !ok {
		this.set[val] = struct{}{}
		this.values = append(this.values, val)
		return !ok
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.set[val]
	if ok {
		delete(this.set, val)
		idx := slices.Index(this.values, val)
		if idx == len(this.values)-1 {
			this.values = this.values[:idx]
		} else {
			this.values = append(this.values[:idx], this.values[idx+1:]...)
		}
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.Intn(len(this.values))
	return this.values[r]
}

// Slow but memory efficient (254ms, 42.57mb)
// type RandomizedSet struct {
// 	set map[int]struct{}
// }
//
// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		set: make(map[int]struct{}),
// 	}
// }
//
// func (this *RandomizedSet) Insert(val int) bool {
// 	_, ok := this.set[val]
// 	this.set[val] = struct{}{}
// 	return !ok
// }
//
// func (this *RandomizedSet) Remove(val int) bool {
// 	if _, ok := this.set[val]; ok {
// 		delete(this.set, val)
// 		return true
// 	}
// 	return false
// }
//
// func (this *RandomizedSet) GetRandom() int {
// 	r := rand.Intn(len(this.set))
// 	i := 0
// 	for v := range this.set {
// 		if r == i {
// 			return v
// 		}
// 		i++
// 	}
// 	return 0
// }
