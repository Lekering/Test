package main

import (
	"unicode"
)

// Оптимизированная версия: O(n) время, O(1) память
// Используем два указателя без создания промежуточного массива
//
//nolint:unused // Решение задачи, используется для обучения
func isPalindrome(s string) bool {
	strRune := []rune(s)

	left, right := 0, len(strRune)-1

	for left < right {
		for left < right && !isVal(strRune[left]) {
			left++
		}
		for right > left && !isVal(strRune[right]) {
			right--
		}
		if unicode.ToLower(strRune[left]) != unicode.ToLower(strRune[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

//nolint:unused // Вспомогательная функция для isPalindrome
func isVal(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}
