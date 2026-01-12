package main

import "fmt"

// Тип для демонстрации типизированного nil
type MyError struct {
	Message string
}

// Реализуем интерфейс error
func (e *MyError) Error() string {
	if e == nil {
		return "nil MyError"
	}
	return e.Message
}

func main() {
	// Функция copy в Go используется для копирования данных из одного среза (slice), массива или байтового среза в другой.
	// Она работает только с срезами или массивами одинакового типа.
	// Сигнатура: n := copy(dst, src)
	// dst и src — это срезы. Копирует min(len(src), len(dst)) элементов.
	// Пример:
	arr := make([]int, 1, 3)

	// ===== ДЕМОНСТРАЦИЯ: Присвоение типа интерфейсу без значения =====

	// 1. Чистый nil интерфейс (нет типа, нет значения)
	var err1 error
	fmt.Printf("err1 == nil: %v (чистый nil, нет типа)\n", err1 == nil)

	// 2. Типизированный nil - интерфейс С ТИПОМ, но БЕЗ ЗНАЧЕНИЯ
	var err2 *MyError = nil       // указатель на тип, равен nil
	var errInterface error = err2 // присваиваем интерфейсу

	fmt.Printf("errInterface == nil: %v ❌ (НЕ nil! содержит тип *MyError)\n", errInterface == nil)
	fmt.Printf("Тип errInterface: %T\n", errInterface)

	// Проверка через рефлексию (чтобы увидеть, что тип есть, но значение nil)
	if errInterface != nil {
		fmt.Printf("  Интерфейс содержит тип, но значение nil!\n")
	}

	// 3. Использование (*Type)(nil) - явное создание типизированного nil
	var typedNil error = (*MyError)(nil)
	fmt.Printf("\ntypedNil == nil: %v ❌ (типизированный nil)\n", typedNil == nil)
	fmt.Printf("Тип typedNil: %T\n", typedNil)

	// 4. Практический пример - когда это важно
	fmt.Println("\n--- Практический пример ---")
	someFunction := func() error {
		return (*MyError)(nil) // возвращаем типизированный nil
	}

	result := someFunction()
	if result != nil {
		fmt.Printf("result != nil, тип: %T ✅\n", result)
		// Правильная проверка типизированного nil
		if myErr, ok := result.(*MyError); ok && myErr == nil {
			fmt.Printf("Это *MyError с nil значением! ✅\n")
		}
	}
	fmt.Printf("arr: %v, len: %d, cap: %d\n", arr, len(arr), cap(arr))

	fmt.Printf("arr: %v\n", arr)

	// Вариант 1: Вернуть новый слайс (правильный способ)
	arr = Appendvalue(arr, 1)
	fmt.Printf("После Appendvalue (с возвратом): %v, len: %d, cap: %d\n", arr, len(arr), cap(arr))

	// Вариант 2: Передать указатель на слайс
	arr2 := make([]int, 1, 3)
	AppendvaluePtr(&arr2, 1)
	fmt.Printf("После AppendvaluePtr (с указателем): %v, len: %d, cap: %d\n", arr2, len(arr2), cap(arr2))

	// Вариант 3: Изменение элементов без append работает!
	arr3 := make([]int, 2, 3)
	arr3[0] = 10
	fmt.Printf("arr3 до: %v\n", arr3)
	ModifyElement(arr3, 0, 999)
	fmt.Printf("arr3 после ModifyElement: %v (элемент изменился!)\n", arr3)

	// Демонстрация: что происходит с append внутри функции
	arr4 := make([]int, 1, 3)
	fmt.Printf("\narr4 до: %v, len: %d, cap: %d, ptr: %p\n", arr4, len(arr4), cap(arr4), arr4)
	appendDemo(arr4, 1)
	fmt.Printf("arr4 после appendDemo: %v, len: %d, cap: %d, ptr: %p\n", arr4, len(arr4), cap(arr4), arr4)
	fmt.Printf("Но внутри функции был новый слайс!\n")

	// Даже если capacity достаточно, изменение локальной переменной не влияет на исходный слайс
	arr5 := make([]int, 1, 3)
	fmt.Printf("\narr5 до: %v, len: %d, cap: %d\n", arr5, len(arr5), cap(arr5))
	appendWithoutRealloc(arr5, 1)
	fmt.Printf("arr5 после (capacity был достаточен, но len не изменился): %v, len: %d\n", arr5, len(arr5))
}

// Правильный способ: вернуть новый слайс
func Appendvalue(arr []int, num int) []int {
	arr = append(arr, num)
	return arr
}

// Альтернатива: использовать указатель на слайс
func AppendvaluePtr(arr *[]int, num int) {
	*arr = append(*arr, num)
}

// Изменение элементов работает, потому что указатель на данные тот же
func ModifyElement(arr []int, index int, value int) {
	if index < len(arr) {
		arr[index] = value
	}
}

// Демонстрация: append создает новый слайс (может быть новый массив, если cap недостаточен)
func appendDemo(arr []int, num int) {
	newArr := append(arr, num)
	fmt.Printf("  Внутри функции после append: %v, len: %d, cap: %d, ptr: %p\n", newArr, len(newArr), cap(newArr), newArr)
}

// Даже если capacity достаточно и массив не переаллоцируется,
// изменение локальной переменной arr не влияет на исходный слайс
func appendWithoutRealloc(arr []int, num int) {
	// arr = append(arr, num) - это создает новую структуру слайса локально
	// Исходный слайс в main не изменится, потому что len и cap - это часть структуры слайса
	arr = append(arr, num)
	fmt.Printf("  Внутри функции: %v, len: %d\n", arr, len(arr))
}
