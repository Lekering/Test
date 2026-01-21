package main

func sValid(s string) bool {
	parenMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}
	for _, r := range s {
		// Если открывающая скобка — добавляем в стек
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else if open, ok := parenMap[r]; ok {
			// r — закрывающая скобка; проверяем соответствие со стеком
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// снимаем верхнюю скобку
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
