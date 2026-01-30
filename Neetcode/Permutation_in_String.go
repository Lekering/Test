package main

func checkInclusion(s1 string, s2 string) bool {
	runeMap1 := make(map[rune]int)

	for _, c := range s1 {
		runeMap1[c]++
	}

	need := len(runeMap1)
	for i := 0; i < len(s2); i++ {
		res := 0
		runeMap2 := make(map[rune]int)
		for j := i; j < len(s2); j++ {
			runeMap2[rune(s2[j])]++
			if runeMap1[rune(s2[j])] < runeMap2[rune(s2[j])] {
				break
			}

			if runeMap2[rune(s2[j])] == runeMap1[rune(s2[j])] {
				res++
			}

			if res == need {
				return true
			}
		}
	}
	return false
}
