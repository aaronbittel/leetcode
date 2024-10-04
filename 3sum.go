package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}

	for i, n := range nums {
		if i > 0 && nums[i-1] == n {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := n + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{n, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return res
}

func threeSumBruteForce(nums []int) [][]int {
	numSet := map[int][]int{}
	for i, n := range nums {
		numSet[n] = append(numSet[n], i)
	}

	fmt.Println(numSet)
	sol := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			r := (nums[i] + nums[j]) * -1
			idxArr, ok := numSet[r]
			if !ok {
				continue
			}
			for _, idx := range idxArr {
				if idx == i || idx == j {
					continue
				}
				candidate := []int{nums[i], nums[j], nums[idx]}
				if isInSolution(sol, candidate) {
					continue
				}
				sol = append(sol, candidate)
			}
		}
	}

	fmt.Println(sol)

	return sol
}

func isInSolution(sols [][]int, candidate []int) bool {
	slices.Sort(candidate)
	for _, s := range sols {
		if slices.Compare(candidate, s) == 0 {
			return true
		}
	}
	return false
}
