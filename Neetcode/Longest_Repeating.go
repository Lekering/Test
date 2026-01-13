package main

func characterReplacement(s string, k int) int {
	coutMap := make(map[byte]int)

	l, res, maxf := 0, 0, 0

	for r := range s {
		coutMap[s[r]]++

		if coutMap[s[r]] > maxf {
			maxf = coutMap[s[r]]
		}

		for (r-l+1)-maxf > k {
			coutMap[s[l]]--
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
