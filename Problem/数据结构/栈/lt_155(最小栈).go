package main

// 题目链接：https://leetcode.cn/problems/min-stack/description/?envType=study-plan-v2&envId=top-100-liked

type MinStack struct {
	data []int
	mind []int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		mind: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.mind) == 0 {
		this.mind = append(this.mind, val)
	} else {
		this.mind = append(this.mind, min(this.mind[len(this.mind)-1], val))
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.mind = this.mind[:len(this.mind)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.mind[len(this.mind)-1]
}
