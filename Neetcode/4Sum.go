package main

import "sort"

func fourSum(nums []int, target int) [][]int {

	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
	}

}
