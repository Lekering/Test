package main

import "sort"

func Sum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := range nums {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			threeSum := a + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
