package main

func removeDuplicates(nums []int) int {

	n := len(nums)
	l, r := 0, 0

	for r < n {
		nums[l] = nums[r]
		for r < n && nums[l] == nums[r] {
			r++
		}
		l++
	}
	return l
}
