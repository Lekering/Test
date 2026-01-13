package main

func lengthOfLongestSubstring(s string) int {
	charmap := make(map[byte]bool)
	l, res := 0, 0

	for r := range s {
		for charmap[s[r]] {
			delete(charmap, s[l])
			l++
		}
		charmap[s[r]] = true
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
