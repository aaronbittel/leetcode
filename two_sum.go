package main

func twoSum2(nums []int, target int) []int {
	set := map[int]int{target - nums[0]: 0}

	for i := 1; i < len(nums); i++ {
		if j, ok := set[nums[i]]; ok {
			return []int{j, i}
		} else {
			set[target-nums[i]] = i
		}
	}
	return nil
}

func twoSumNaive(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
