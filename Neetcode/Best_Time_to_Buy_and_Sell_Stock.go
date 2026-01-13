package main

func BestTime(prices []int) int {
	l, r := 0, 1
	maxProfit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			prorit := prices[r] - prices[l]
			if maxProfit < prorit {
				maxProfit = prorit
			}
		} else {
			l = r
		}
		r++
	}
	return maxProfit
}
