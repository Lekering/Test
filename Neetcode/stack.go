package main

func sValid(s string) bool {
	stack := []rune{}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (r == ')' && last != '(') ||
				(r == '}' && last != '{') ||
				(r == ']' && last != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
