package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i, n := range nums2 {
			nums1[i] = n
		}
		return
	}

	i, j := 0, 0
	for j < n {
		if i < m && nums1[i] <= nums2[j] {
			i++
			continue
		}
		for k := len(nums1) - 1; k > i; k-- {
			nums1[k] = nums1[k-1]
		}
		nums1[i] = nums2[j]
		j++
		m++
		fmt.Println(nums1)
	}
}
