package main

import "fmt"

func main() {
	// true, false, true, 2, true, false, 2
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Remove(2))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet.Insert(1))
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.Remove(1))
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.Insert(4))
	fmt.Println(randomizedSet.Insert(8))
	fmt.Println(randomizedSet.Insert(122))
	fmt.Println(randomizedSet.Insert(122))
	fmt.Println(randomizedSet.Insert(122))
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.Remove(4))
	fmt.Println(randomizedSet)
	// fmt.Println(randomizedSet.Remove(2), "==?", false)
	// fmt.Println(randomizedSet.Insert(1), "==?", true)
	// fmt.Println(randomizedSet.Insert(2), "==?", true)
	// fmt.Println(randomizedSet.GetRandom(), "==?", 2)
	// fmt.Println(randomizedSet.Remove(1), "==?", true)
	// fmt.Println(randomizedSet.Insert(2), "==?", false)
	// fmt.Println(randomizedSet.GetRandom(), "==?", 2)
}
