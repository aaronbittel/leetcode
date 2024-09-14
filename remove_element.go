package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for range len(nums) {
		if nums[i] != val {
			fmt.Println("NOT EQUAL", nums, i)
			i++
			continue
		}

		fmt.Println("EQUAL BEFORE", nums, i)
		for j := i; j < len(nums)-1-i; j++ {
			nums[j] = nums[j+1]
		}
	}
	return i
}

func removeElementNotInplace(nums []int, val int) int {
	arr := []int{}
	for _, n := range nums {
		if n != val {
			arr = append(arr, n)
			continue
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < len(arr) {
			nums[i] = arr[i]
		} else {
			nums[i] = 0
		}
	}
	return len(arr)
}
