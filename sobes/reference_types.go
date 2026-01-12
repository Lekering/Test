package main

import "fmt"

func main() {
	fmt.Println("=== ДЕМОНСТРАЦИЯ ТИПОВ С УКАЗАТЕЛЯМИ ВНУТРИ ===\n")

	// 1. СЛАЙСЫ (Slices)
	fmt.Println("1. СЛАЙСЫ:")
	demonstrateSlice()

	// 2. МАПЫ (Maps)
	fmt.Println("\n2. МАПЫ:")
	demonstrateMap()

	// 3. КАНАЛЫ (Channels)
	fmt.Println("\n3. КАНАЛЫ:")
	demonstrateChannel()

	// 4. ИНТЕРФЕЙСЫ (Interfaces)
	fmt.Println("\n4. ИНТЕРФЕЙСЫ:")
	demonstrateInterface()

	// 5. УКАЗАТЕЛИ (Pointers) - очевидно
	fmt.Println("\n5. УКАЗАТЕЛИ:")
	demonstratePointer()

	// 6. ФУНКЦИИ (Functions) - могут содержать замыкания
	fmt.Println("\n6. ФУНКЦИИ:")
	demonstrateFunction()

	// 7. СТРУКТУРЫ со слайсами/мапами внутри
	fmt.Println("\n7. СТРУКТУРЫ со слайсами/мапами:")
	demonstrateStruct()
}

// ========== 1. СЛАЙСЫ ==========
func demonstrateSlice() {
	s := []int{1, 2, 3}
	fmt.Printf("  Исходный слайс: %v\n", s)

	// Изменение элементов работает (меняем данные по указателю)
	modifySlice(s, 0, 999)
	fmt.Printf("  После изменения элемента: %v ✅ (элемент изменился)\n", s)

	// Но присваивание нового слайса не работает
	reassignSlice(s, []int{10, 20, 30})
	fmt.Printf("  После переприсваивания: %v ❌ (не изменился, т.к. это новая локальная копия структуры)\n", s)
}

func modifySlice(s []int, index int, value int) {
	if index < len(s) {
		s[index] = value // ✅ Работает - меняем данные по указателю
	}
}

func reassignSlice(s []int, newSlice []int) {
	s = newSlice // ❌ Не работает - это локальная копия структуры слайса
}

// ========== 2. МАПЫ ==========
func demonstrateMap() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("  Исходная мапа: %v\n", m)

	// Изменение элементов мапы работает
	modifyMap(m, "a", 999)
	fmt.Printf("  После изменения значения: %v ✅ (значение изменилось)\n", m)

	// Добавление элементов работает
	addToMap(m, "c", 3)
	fmt.Printf("  После добавления элемента: %v ✅ (элемент добавился)\n", m)

	// Но переприсваивание не работает
	reassignMap(m, map[string]int{"x": 100})
	fmt.Printf("  После переприсваивания: %v ❌ (не изменилась)\n", m)
}

func modifyMap(m map[string]int, key string, value int) {
	m[key] = value // ✅ Работает - мапа содержит указатель на данные
}

func addToMap(m map[string]int, key string, value int) {
	m[key] = value // ✅ Работает
}

func reassignMap(m map[string]int, newMap map[string]int) {
	m = newMap // ❌ Не работает - локальная копия указателя на мапу
}

// ========== 3. КАНАЛЫ ==========
func demonstrateChannel() {
	ch := make(chan int, 3)

	// Каналы передаются по значению (копия указателя на канал)
	// Но это та же самая структура канала, поэтому отправка/прием работают
	sendToChannel(ch, 42)

	go func() {
		value := <-ch
		fmt.Printf("  Значение из канала: %d ✅ (канал работает)\n", value)
	}()

	// Но переприсваивание не работает
	reassignChannel(ch, make(chan int))
	fmt.Printf("  После переприсваивания канал все еще работает: %v ✅\n", len(ch) == 1)
}

func sendToChannel(ch chan int, value int) {
	ch <- value // ✅ Работает - это тот же канал
}

func reassignChannel(ch chan int, newCh chan int) {
	ch = newCh // ❌ Не влияет на исходный канал
}

// ========== 4. ИНТЕРФЕЙСЫ ==========
func demonstrateInterface() {
	var i interface{} = 42
	fmt.Printf("  Исходный интерфейс: %v (тип: %T)\n", i, i)

	// Интерфейсы содержат указатель на данные и таблицу методов
	modifyInterface(&i, "hello")
	fmt.Printf("  После изменения: %v (тип: %T) ✅\n", i, i)
}

func modifyInterface(i *interface{}, value interface{}) {
	*i = value // Работает, т.к. передали указатель
}

// ========== 5. УКАЗАТЕЛИ ==========
func demonstratePointer() {
	x := 42
	p := &x
	fmt.Printf("  Исходное значение через указатель: %d\n", *p)

	modifyPointer(p, 999)
	fmt.Printf("  После изменения через указатель: %d ✅ (значение изменилось)\n", *p)
}

func modifyPointer(p *int, value int) {
	*p = value // ✅ Очевидно работает
}

// ========== 6. ФУНКЦИИ ==========
func demonstrateFunction() {
	counter := 0
	inc := func() {
		counter++ // Замыкание захватывает переменную
	}

	fmt.Printf("  Счетчик до: %d\n", counter)
	callFunction(inc)
	fmt.Printf("  Счетчик после: %d ✅ (функция с замыканием работает)\n", counter)
}

func callFunction(f func()) {
	f() // Вызов функции с замыканием
}

// ========== 7. СТРУКТУРЫ ==========
type Container struct {
	Slice []int
	Map   map[string]int
	Value int
}

func demonstrateStruct() {
	c := Container{
		Slice: []int{1, 2, 3},
		Map:   map[string]int{"a": 1},
		Value: 42,
	}

	fmt.Printf("  Исходная структура: Slice=%v, Map=%v, Value=%d\n", c.Slice, c.Map, c.Value)

	modifyStruct(c)
	fmt.Printf("  После modifyStruct: Slice=%v, Map=%v, Value=%d\n", c.Slice, c.Map, c.Value)
	fmt.Printf("    Slice: %v ✅ (изменился - слайс внутри)\n", c.Slice)
	fmt.Printf("    Map: %v ✅ (изменился - мапа внутри)\n", c.Map)
	fmt.Printf("    Value: %d ❌ (НЕ изменился - обычное поле передается по значению)\n", c.Value)
}

func modifyStruct(c Container) {
	c.Slice[0] = 999        // ✅ Работает - слайс содержит указатель
	c.Map["a"] = 999        // ✅ Работает - мапа содержит указатель
	c.Value = 999           // ❌ Не работает - обычное поле
	c.Slice = []int{10, 20} // ❌ Не работает - переприсваивание слайса в структуре
}
