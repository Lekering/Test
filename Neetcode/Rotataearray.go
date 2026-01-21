package main

// Вариант эффективнее: in-place reverse, O(n) time, O(1) space
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || k%n == 0 {
		return
	}
	k = k % n // на случай если k > n
	reverse := func(arr []int, l, r int) {
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}
