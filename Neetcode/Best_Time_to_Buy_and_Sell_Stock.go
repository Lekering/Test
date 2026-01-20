package main

func BestTime(prices []int) int {
	l, r := 0, 1
	maxprofit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if maxprofit < profit {
				maxprofit = profit
			}
		} else {
			l = r
		}
		r++
	}
	return maxprofit
}
