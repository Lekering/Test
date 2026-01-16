package main

func isValid(s string) bool {
	stack := []rune{}
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if math, found := hash[char]; found {
			if len(stack) > 0 && math == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
