package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sunTarget := numbers[left] + numbers[right]
		if sunTarget == target {
			return []int{left + 1, right + 1}
		}
		if sunTarget > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
