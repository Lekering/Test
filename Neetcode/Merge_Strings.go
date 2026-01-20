package main

import "strings"

func mergeAlternately(w1, w2 string) string {
	var res strings.Builder

	i, j := 0, 0
	n, m := len(w1), len(w2)

	for i < n || j < m {
		if i < n {
			res.WriteByte(w1[i])
			i++
		}
		if j < m {
			res.WriteByte(w2[j])
			j++
		}
	}
	return res.String()
}
