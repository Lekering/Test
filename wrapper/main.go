package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Statistic struct {
	id   int
	step int
}

type Result struct {
	id   []int
	step int
}

var stat = [][]Statistic{
	{{id: 1, step: 600}, {id: 2, step: 300}},
	{{id: 1, step: 200}, {id: 2, step: 500}},
}

// подспрась
// АНАЛИЗ СЛОЖНОСТИ ПО ПАМЯТИ:
// Входные данные: statistic [][]Statistic - передается по ссылке на внутренний массив (O(1) дополнительной памяти)
//
// Внутри функции:
// 1. peopleCount map[int]int - хранит количество дней для каждого уникального ID
// 2. totalStep map[int]int - хранит сумму шагов для каждого уникального ID
// 3. champions []int - слайс для хранения чемпионов
//
// Обозначения:
//
//	N = общее количество записей (Statistic) во всех днях
//	D = количество дней (len(statistic))
//	U = количество уникальных ID людей (может быть <= N)
//
// Сложность по памяти: O(U)
//   - peopleCount: максимум U записей = O(U)
//   - totalStep: максимум U записей = O(U)
//   - champions: в худшем случае все U людей могут быть чемпионами = O(U)
//   - Итого: O(U) + O(U) + O(U) = O(U)
//
// В лучшем случае (только 1 чемпион): O(U) для карт + O(1) для champions = O(U)
// В худшем случае (все U людей - чемпионы): O(U) для карт + O(U) для champions = O(U)
//
// Примечание: U <= N, так что можно также сказать O(N) в терминах входных данных
func getChampions(statistic [][]Statistic) Result {
	if len(statistic) == 0 {
		return Result{}
	}

	allDays := len(statistic)
	peopleCount := make(map[int]int) // O(U) памяти, где U = количество уникальных ID
	totalStep := make(map[int]int)   // O(U) памяти

	// Проходим по всем записям: временная сложность O(N), память уже выделена выше
	for _, day := range statistic {
		for _, res := range day {
			peopleCount[res.id]++         // Добавление/обновление в карте: O(1) амортизированно
			totalStep[res.id] += res.step // O(1) амортизированно
		}
	}

	maxStep := 0
	var champions []int // Начальная длина 0, будет расти до O(U) в худшем случае

	// Проходим по карте: временная сложность O(U)
	for id, count := range peopleCount {
		if count == allDays {
			if totalStep[id] > maxStep {
				maxStep = totalStep[id]
				// champions = []int{id} - создаём новый срез, старый сборщик мусора может удалить
				// В худшем случае это может происходить до U раз, но каждый раз старый слайс освобождается
				champions = []int{id}
			} else if totalStep[id] == maxStep {
				// append может реаллоцировать память, но амортизированно O(1) на операцию
				champions = append(champions, id) // O(1) амортизированно
			}
		}
	}

	return Result{
		id:   champions, // O(U) в худшем случае, O(1) в лучшем
		step: maxStep,
	}
}

// analyzeMemoryComplexity демонстрирует анализ памяти для понимания
func analyzeMemoryComplexity(statistic [][]Statistic) {
	fmt.Println("=== АНАЛИЗ СЛОЖНОСТИ ПО ПАМЯТИ ===")

	// Подсчитываем параметры
	totalRecords := 0
	uniqueIDs := make(map[int]bool)

	for _, day := range statistic {
		for _, res := range day {
			totalRecords++
			uniqueIDs[res.id] = true
		}
	}

	N := totalRecords   // Общее количество записей
	U := len(uniqueIDs) // Количество уникальных ID
	D := len(statistic) // Количество дней

	fmt.Printf("N (всего записей) = %d\n", N)
	fmt.Printf("U (уникальных ID) = %d\n", U)
	fmt.Printf("D (дней) = %d\n", D)
	fmt.Println()

	fmt.Println("Использование памяти:")
	fmt.Printf("  peopleCount map: O(U) = O(%d) записей\n", U)
	fmt.Printf("  totalStep map:   O(U) = O(%d) записей\n", U)
	fmt.Printf("  champions slice: O(U) = O(%d) в худшем случае\n", U)
	fmt.Println()
	fmt.Printf("ИТОГО: O(U) = O(%d)\n", U)
	fmt.Printf("(В терминах входных данных: O(N) = O(%d))\n", N)
	fmt.Println()

	if U == 1 {
		fmt.Println("Лучший случай: 1 уникальный ID")
		fmt.Println("  Память для champions: O(1)")
		fmt.Printf("  Итого: O(%d) ≈ O(1)\n", U)
	} else if U <= N/2 {
		fmt.Printf("Средний случай: U < N (меньше чем половина записей)\n")
		fmt.Printf("  Память: O(%d) < O(%d)\n", U, N)
	}

	// Реальный результат для демонстрации
	result := getChampions(statistic)
	fmt.Println()
	fmt.Printf("Результат: champions = %v (длина: %d), maxStep = %d\n",
		result.id, len(result.id), result.step)
}

func main() {
	result := getChampions(stat)
	fmt.Printf("getChampions(stat): %v\n", result)
	fmt.Println()
	analyzeMemoryComplexity(stat)
}

func LongWork() int {
	res := rand.Intn(500)
	time.Sleep(time.Second * 3)
	return res
}

func WrapperWithContext(ctx context.Context, task func() int) (int, error) {
	chanresout := make(chan int)

	go func() {

		res := LongWork()

		select {
		case chanresout <- res:
		default:
		}
	}()

	select {
	case result := <-chanresout:
		return result, nil
	case <-ctx.Done():
		return -1, errors.New("SUCK")

	}
}

func Wrapper(timeout time.Duration, task func() int) (*int, error) {
	chanresout := make(chan int)
	timechan := time.After(timeout)

	go func() {

		res := task()

		select {
		case chanresout <- res:
		default:
		}
	}()

	select {
	case result := <-chanresout:
		return &result, nil
	case <-timechan:
		return nil, errors.New("SUCK")

	}
}
