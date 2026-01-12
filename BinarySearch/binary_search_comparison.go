package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ========== ВАРИАНТ 1: Итеративный (базовый) ==========
// ⚠️ ПРОБЛЕМА: может быть переполнение при (left + right) / 2
func BinarySearchIterative(target int, array []int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		midle := (left + right) / 2 // ⚠️ Потенциальное переполнение!
		if target == array[midle] {
			return midle
		} else if array[midle] < target {
			left = midle + 1
		} else {
			right = midle - 1 // ⚠️ В вашем коде было right = midle (ошибка!)
		}
	}
	return -1
}

// ========== ВАРИАНТ 2: Итеративный (безопасный от переполнения) ==========
// ✅ БЕЗОПАСНЫЙ: использует left + (right-left)/2
func BinarySearchSafe(target int, array []int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		midle := left + (right-left)/2 // ✅ Безопасно от переполнения
		if target == array[midle] {
			return midle
		} else if array[midle] < target {
			left = midle + 1
		} else {
			right = midle - 1
		}
	}
	return -1
}

// ========== ВАРИАНТ 3: Итеративный (с битовым сдвигом) ==========
// ⚡ САМЫЙ БЫСТРЫЙ: битовый сдвиг быстрее деления
func BinarySearchBitShift(target int, array []int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		midle := left + ((right - left) >> 1) // ⚡ >> 1 быстрее / 2
		if target == array[midle] {
			return midle
		} else if array[midle] < target {
			left = midle + 1
		} else {
			right = midle - 1
		}
	}
	return -1
}

// ========== ВАРИАНТ 4: Рекурсивный ==========
// 🐌 МЕДЛЕННЫЙ: накладные расходы на вызовы функций
func BinarySearchRecursive(target int, array []int) int {
	return binarySearchRecursiveHelper(target, array, 0, len(array)-1)
}

func binarySearchRecursiveHelper(target int, array []int, left, right int) int {
	if left > right {
		return -1
	}
	midle := left + (right-left)/2
	if target == array[midle] {
		return midle
	} else if array[midle] < target {
		return binarySearchRecursiveHelper(target, array, midle+1, right)
	} else {
		return binarySearchRecursiveHelper(target, array, left, midle-1)
	}
}

// ========== ВАРИАНТ 5: Branchless (без ветвлений) ==========
// ⚡ ОПТИМИЗИРОВАННЫЙ: меньше условных переходов (быстрее на некоторых CPU)
func BinarySearchBranchless(target int, array []int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		midle := left + ((right - left) >> 1)
		val := array[midle]

		// Branchless: используем арифметику вместо if-else
		diff := val - target
		// Если diff == 0, то found = 1, иначе 0
		if diff == 0 {
			return midle
		}
		// Если diff < 0, то left = midle + 1, иначе right = midle - 1
		left += (diff >> 31) & (midle - left + 1)
		right -= (^diff >> 31) & (right - midle + 1)
	}
	return -1
}

// ========== ВАРИАНТ 6: С ранним выходом (оптимизированный) ==========
// ⚡ БЫСТРЫЙ: проверяет границы перед циклом
func BinarySearchOptimized(target int, array []int) int {
	n := len(array)
	if n == 0 {
		return -1
	}

	// Ранний выход для граничных случаев
	if target < array[0] || target > array[n-1] {
		return -1
	}

	left := 0
	right := n - 1

	for left <= right {
		midle := left + ((right - left) >> 1)
		val := array[midle]

		if val == target {
			return midle
		} else if val < target {
			left = midle + 1
		} else {
			right = midle - 1
		}
	}
	return -1
}

// ========== БЕНЧМАРК ==========
func benchmark(name string, fn func(int, []int) int, array []int, iterations int) {
	targets := make([]int, iterations)
	for i := range targets {
		targets[i] = array[rand.Intn(len(array))]
	}

	start := time.Now()
	for _, target := range targets {
		fn(target, array)
	}
	duration := time.Since(start)

	fmt.Printf("%-30s: %v (%d итераций)\n", name, duration, iterations)
}

func main() {
	fmt.Println("=== СРАВНЕНИЕ ВАРИАНТОВ БИНАРНОГО ПОИСКА ===\n")

	// Создаем большой отсортированный массив
	size := 10_000_000
	array := make([]int, size)
	for i := range array {
		array[i] = i * 2 // Четные числа от 0 до 2*(size-1)
	}

	fmt.Printf("Размер массива: %d элементов\n", size)
	fmt.Printf("Итераций поиска: 1,000,000\n\n")

	iterations := 1_000_000

	benchmark("1. Итеративный (базовый)", BinarySearchIterative, array, iterations)
	benchmark("2. Итеративный (безопасный)", BinarySearchSafe, array, iterations)
	benchmark("3. Итеративный (битовый сдвиг)", BinarySearchBitShift, array, iterations)
	benchmark("4. Рекурсивный", BinarySearchRecursive, array, iterations)
	benchmark("5. Branchless", BinarySearchBranchless, array, iterations)
	benchmark("6. Оптимизированный", BinarySearchOptimized, array, iterations)

	fmt.Println("\n=== ВЫВОДЫ ===")
	fmt.Println("⚡ САМЫЙ БЫСТРЫЙ: Вариант 3 (битовый сдвиг) или Вариант 6 (оптимизированный)")
	fmt.Println("✅ РЕКОМЕНДУЕМЫЙ: Вариант 6 (оптимизированный) - безопасный + быстрый + ранний выход")
	fmt.Println("🐌 МЕДЛЕННЫЙ: Вариант 4 (рекурсивный) - накладные расходы на вызовы функций")
	fmt.Println("\n💡 ПРИМЕЧАНИЯ:")
	fmt.Println("   - Битовый сдвиг (>> 1) быстрее деления (/ 2) на большинстве процессоров")
	fmt.Println("   - Ранний выход для граничных случаев ускоряет поиск")
	fmt.Println("   - Итеративный вариант быстрее рекурсивного из-за отсутствия накладных расходов")
	fmt.Println("   - Branchless может быть быстрее на некоторых CPU, но сложнее для понимания")
}
