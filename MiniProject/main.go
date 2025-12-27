package main

import "fmt"

type cmd func(in, out chan any)

func generating() chan any {
	out := make(chan any)

	go func() {
		defer close(out) // defer гарантирует закрытие после завершения цикла
		for i := range 20 {
			out <- i
		}
	}()
	return out
}

// SumNumber принимает множитель как аргумент и возвращает функцию типа cmd
func SumNumber(multiplier int) cmd {
	return func(in, out chan any) {
		defer close(out)
		for val := range in {
			if num, ok := val.(int); ok {
				out <- num * multiplier
			} else {
				out <- val
			}
		}
	}
}

// FilterEven принимает дополнительный аргумент (можно использовать для другой логики)
func FilterEven(in, out chan any) {
	defer close(out)
	for val := range in {
		if num, ok := val.(int); ok {
			if num%2 == 0 {
				out <- num
			}
		}
	}
}

// FilterByStrings принимает массив строк для фильтрации
func FilterByStrings(allowedStrings []string) cmd {
	return func(in, out chan any) {
		defer close(out)
		// Создаем map для быстрого поиска
		allowed := make(map[string]bool)
		for _, s := range allowedStrings {
			allowed[s] = true
		}

		for val := range in {
			if str, ok := val.(string); ok {
				if allowed[str] {
					out <- str
				}
			} else {
				out <- val
			}
		}
	}
}

// AddConstant принимает константу как аргумент
func AddConstant(constant int) cmd {
	return func(in, out chan any) {
		defer close(out)
		for val := range in {
			if num, ok := val.(int); ok {
				out <- num + constant
			} else {
				out <- val
			}
		}
	}
}

// TransformStrings принимает массив строк-замен (из -> в)
func TransformStrings(replacements map[string]string) cmd {
	return func(in, out chan any) {
		defer close(out)
		for val := range in {
			switch str := val.(type) {
			case string:
				if replacement, exists := replacements[str]; exists {
					out <- replacement
				} else {
					out <- str
				}
			default:
				out <- val
			}
		}
	}
}

// GenerateStrings генерирует строки из массива
func GenerateStrings(strings []string) chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		for _, s := range strings {
			out <- s
		}
	}()
	return out
}

// Pipeline принимает начальный канал и команды
func Pipeline(startChan chan any, comand ...cmd) <-chan any {
	// Связываем этапы pipeline: выход предыдущего становится входом следующего
	currentIn := startChan
	for _, c := range comand {
		out := make(chan any)
		go c(currentIn, out)
		currentIn = out // Выход становится входом для следующего этапа
	}

	return currentIn // Возвращаем последний выходной канал
}

func main() {
	fmt.Println("=== Пример 1: Работа с числами ===")
	// Используем функции с аргументами через замыкания
	result1 := Pipeline(
		generating(),
		SumNumber(3),     // умножаем на 3
		FilterEven,       // фильтруем четные
		AddConstant(100), // добавляем 100
	)

	for val := range result1 {
		fmt.Println("Число:", val)
	}

	fmt.Println("\n=== Пример 2: Работа со строками ===")
	// Генерируем строки из массива
	strings := []string{"apple", "banana", "apple", "orange", "banana", "grape"}

	// Фильтруем только разрешенные строки
	allowed := []string{"apple", "banana", "orange"}

	// Заменяем некоторые строки
	replacements := map[string]string{
		"apple":  "яблоко",
		"banana": "банан",
		"orange": "апельсин",
	}

	result2 := Pipeline(
		GenerateStrings(strings),       // генерируем строки
		FilterByStrings(allowed),       // фильтруем по массиву строк
		TransformStrings(replacements), // заменяем строки
	)

	for val := range result2 {
		fmt.Println("Строка:", val)
	}
}
