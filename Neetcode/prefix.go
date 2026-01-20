package main

func Prefix(arr []int) []int {

	res := make([]int, len(arr)+1)

	for i := range arr {
		res[i+1] = arr[i] + res[i]
	}
	return res
}

func Findnum(arr []int, l, r int) int {

	res := Prefix(arr)

	return res[r+1] - res[l]
}
