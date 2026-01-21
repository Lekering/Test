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
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1

			sum := nums[i] + nums[j] + nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
				left++
				right--
			}

		}
	}
}
