package main

type MinStack struct {
	stack  []int
	mStack []int
	size   int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		mStack: make([]int, 0),
		size:   0,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.size == 0 {
		this.mStack = append(this.mStack, val)
	} else {
		minVal := minInt(this.mStack[this.size-1], val)
		this.mStack = append(this.mStack, minVal)
	}
	this.size++
}

func (this *MinStack) Pop() {
	this.size--
	this.stack = this.stack[:this.size]
	this.mStack = this.mStack[:this.size]
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.mStack[this.size-1]
}
