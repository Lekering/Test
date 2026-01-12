package main

type MinStack struct {
	minstack []int
	stack    []int
}

func Constructor() MinStack {
	return MinStack{
		minstack: []int{},
		stack:    []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, val)
		this.minstack = append(this.minstack, val)
	} else {
		minvalue := this.minstack[len(this.minstack)-1]
		this.stack = append(this.stack, val)
		if val < minvalue {
			this.minstack = append(this.minstack, val)
		} else {
			this.minstack = append(this.minstack, minvalue)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	lastvalue := len(this.stack) - 1
	this.minstack = this.minstack[:lastvalue]
	this.stack = this.stack[:lastvalue]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	top := this.stack[len(this.stack)-1]
	return top
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	min := this.minstack[len(this.minstack)-1]
	return min
}
