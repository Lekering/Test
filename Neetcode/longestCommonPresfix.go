package main

func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}

	s := strs[0]

	for _, str := range strs {
		if len(str) < len(s) {
			s = str
		}
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		for _, str := range strs {
			if str[i] != c {
				return res
			}
		}
		res += string(c)
	}
	return res
}
