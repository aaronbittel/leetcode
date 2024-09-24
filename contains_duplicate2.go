package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	hashSet := make(map[int]struct{})

	left := 0
	for right, n := range nums {
		if right-left > k {
			delete(hashSet, nums[left])
			left++
		}
		if _, ok := hashSet[n]; ok {
			return true
		}
		hashSet[n] = struct{}{}
	}

	return false
}
