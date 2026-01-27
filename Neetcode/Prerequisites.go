package main

func sortColors(nums []int) {
	res := make([]int, 3)

	for _, num := range nums {
		res[num]++
	}
	index := 0
	for i := range res {
		for res[i] > 0 {
			res[i]--
			nums[index] = i
			index++
		}
	}
}
