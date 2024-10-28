package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if target > nums[r] {
		return r + 1
	}

	for l < r {
		mid := l + ((r - l) / 2)

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if nums[l] < target {
		return l + 1
	}

	return l
}
