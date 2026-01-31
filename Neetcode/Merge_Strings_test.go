package main

import "testing"

// Здесь нужно реализовать тесты для функции mergeAlternately.
// 1. Заполнить слайс tests валидными тест-кейсами: задать значения w1, w2 и ожидаемый результат want.
// 2. Внутри каждого теста сравнить результат вызова mergeAlternately с ожидаемым значением want с помощью if got != tt.want.
// 3. Если результат не совпадает с ожидаемым, вызвать t.Errorf для вывода ошибки.

func Test_mergeAlternately(t *testing.T) {
	tests := []struct {
		name string
		w1   string
		w2   string
		want string
	}{
		{
			name: "одинаковая длина",
			w1:   "abc",
			w2:   "xyz",
			want: "axbycz",
		},
		{
			name: "w1 длиннее",
			w1:   "abcd",
			w2:   "xy",
			want: "axbycd",
		},
		{
			name: "w2 длиннее",
			w1:   "ab",
			w2:   "wxyz",
			want: "awbxyz",
		},
		{
			name: "одна строка пустая",
			w1:   "",
			w2:   "test",
			want: "test",
		},
		{
			name: "обе строки пустые",
			w1:   "",
			w2:   "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.w1, tt.w2)
			if got != tt.want {
				t.Errorf("mergeAlternately(%q, %q) = %v, want %v", tt.w1, tt.w2, got, tt.want)
			}
		})
	}
}
