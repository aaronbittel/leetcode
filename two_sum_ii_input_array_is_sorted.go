package main

// 3 ms
func twoSum(numbers []int, target int) []int {
	delta := map[int]int{}

	for i, v := range numbers {
		if j, ok := delta[v]; ok {
			return []int{j + 1, i + 1}
		}
		delta[target-v] = i
	}

	return []int{}
}

// binary search like: right jumps to the middle if sum > target, right jumps to
// left side else right side, ... and so on
// 11ms
func twoSumBinary(numbers []int, target int) []int {
	var sum int
	for left := 0; left < len(numbers); left++ {
		low, high := left+1, len(numbers)-1
		for low <= high {
			mid := low + (high-low)/2
			sum = numbers[left] + numbers[mid]
			if sum == target {
				return []int{left + 1, mid + 1}
			}
			if sum > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{}
}

// 619ms
func twoSumSlow(numbers []int, target int) []int {
	left, right := 0, 1
	var sum int
	for {
		sum = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			left++
			right = left + 1
			continue
		}
		right++
		if right >= len(numbers) {
			left++
			right = left + 1
		}
	}
}
