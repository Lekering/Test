package main

func validPalindrome(s string) bool {

	validChek := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return validChek(l+1, r) || validChek(l, r-1)
		}
		l++
		r--
	}
	return true
}
