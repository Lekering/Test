package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	pointer := 1
	zeroCount := 0

	for i := range nums {
		if nums[i] != 0 {
			pointer *= nums[i]
		} else {
			zeroCount++
		}
	}
	if zeroCount > 1 {
		return res
	}

	for i, v := range nums {
		if zeroCount > 0 {
			if v != 0 {
				res[i] = 0
			} else {
				res[i] = pointer
			}
		} else {
			res[i] = pointer / v
		}
	}
	fmt.Println(res)
	return res
}

func PrefAndPost(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := range nums {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	fmt.Println(res)
	return res
}
