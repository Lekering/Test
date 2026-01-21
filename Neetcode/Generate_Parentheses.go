package main

// Эта функция генерирует все возможные корректные строки из n пар скобок.
// Вот что происходит поэтапно:
//
//nolint:unused // Решение задачи, используется для обучения
func generateParenthesis(n int) []string {
	var solve func(open, close int, temp string)
	ans := []string{}

	solve = func(open, close int, temp string) {
		if len(temp) == 2*n {
			ans = append(ans, temp)
			return
		}

		if open > 0 {
			solve(open-1, close, temp+"(")
		}
		if close > open {
			solve(open, close-1, temp+")")
		}
	}

	solve(n, n, "")
	return ans
}
