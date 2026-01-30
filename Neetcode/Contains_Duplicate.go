package main

func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]struct{})
	l := 0

	for r := range nums {

		if r-l > k {
			delete(window, nums[l])
			l++
		}

		if _, ok := window[nums[r]]; ok {
			return true
		}

		window[nums[r]] = struct{}{}
	}
	return false
}
