package main

type Stack struct {
	elems []int
	size  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		elems: make([]int, 0, cap),
		size:  0,
	}
}

func (stk *Stack) Len() int {
	return stk.size
}

func (stk *Stack) Top() int {
	if stk.size == 0 {
		panic("empty stack")
	}
	return stk.elems[stk.size-1]
}

func (stk *Stack) Pop() int {
	if stk.size == 0 {
		panic("empty stack")
	}

	elem := stk.elems[stk.size-1]
	stk.elems = stk.elems[:stk.size-1]
	stk.size--
	return elem
}

func (stk *Stack) Push(x int) {
	stk.elems = append(stk.elems, x)
	stk.size++
}
