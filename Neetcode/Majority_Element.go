package main

func majorityElement(nums []int) int {
	res, count := 0, len(nums)/2

	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num]++
		if numMap[num] > count {
			res = num
			count = numMap[num]
		}
	}
	return res
}
