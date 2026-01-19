package main

import (
	"fmt"
	"sync"
)

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))

	arr := []int{1, 2, 3, 4, 5}
	arr1 := arr[0:5]
	arr2 := arr[2:4]
	fmt.Println(arr1, arr2)
}

func Revers(array []int) {
	f, s := 0, len(array)-1

	for f < s {
		array[f], array[s] = array[s], array[f]
		f++
		s--
	}
	fmt.Println(array)
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

// Функция letterCombinations генерирует все возможные комбинации букв,
// соответствующих последовательности цифр, как на телефонной клавиатуре.
func letterCombinations(digits string) []string {
	// Шаг 1: Проверяем, была ли передана пустая строка. Если да, возвращаем nil.
	if digits == "" {
		return nil
	}

	// Шаг 2: Создаём отображение цифр к возможным буквам, как на телефонной клавиатуре.
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// Шаг 3: Создаём срез для хранения всех комбинаций результата.
	var res []string

	// Шаг 4: Описываем функцию backtrack (обратного хода), которая будет рекурсивно собирать комбинации.
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		// Базовый случай: если собрали путь длины digits, добавляем результат.
		if index == len(digits) {
			res = append(res, path)
			return
		}
		// Шаг 5: Получаем возможные буквы для текущей цифры.
		letters := digitToLetters[digits[index]]

		// Шаг 6: Перебираем все буквы, добавляем к текущему пути и вызываем backtrack для следующей позиции.
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, path+string(letters[i]))
		}
	}

	// Шаг 7: Запускаем обратный ход с начальной позиции и пустой строкой.
	backtrack(0, "")

	// Возвращаем все возможные комбинации.
	return res
}
