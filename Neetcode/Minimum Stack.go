package main

type MinStack struct {
	minstack []int
	stack    []int
}

func NewMinStack() MinStack {
	return MinStack{
		minstack: []int{},
		stack:    []int{},
	}
}

func (ms *MinStack) Push(val int) {
	if len(ms.stack) == 0 {
		ms.stack = append(ms.stack, val)
		ms.minstack = append(ms.minstack, val)
	} else {
		minvalue := ms.minstack[len(ms.minstack)-1]
		ms.stack = append(ms.stack, val)
		if val < minvalue {
			ms.minstack = append(ms.minstack, val)
		} else {
			ms.minstack = append(ms.minstack, minvalue)
		}
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	lastvalue := len(ms.stack) - 1
	ms.minstack = ms.minstack[:lastvalue]
	ms.stack = ms.stack[:lastvalue]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	top := ms.stack[len(ms.stack)-1]
	return top
}

func (ms *MinStack) GetMin() int {
	if len(ms.stack) == 0 {
		return 0
	}
	min := ms.minstack[len(ms.minstack)-1]
	return min
}
