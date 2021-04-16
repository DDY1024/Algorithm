package main

import "fmt"

func lowBit(x int) int {
	return x & (-x)
}

func update(idx, n, c int, tree []int) {
	for i := idx; i <= n; i += lowBit(i) {
		tree[i] += c
	}
}

func getSum(idx int, tree []int) int {
	ret := 0
	for i := idx; i > 0; i -= lowBit(i) {
		ret += tree[i]
	}
	return ret
}

func findK(k int, arr []int) (int, int) {
	l, r, ret := 1, 100000, -1
	for l <= r {
		mid := l + (r-l)>>1
		if getSum(mid, arr) >= k {
			ret, r = mid, mid-1
		} else {
			l = mid + 1
		}
	}
	return ret, getSum(ret, arr) - k
}

type MKAverage struct {
	m, k, div   int
	front, rear int
	ele         []int
	sum         []int
	cnt         []int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:   m,
		k:   k,
		div: m - 2*k,
		ele: make([]int, 100000+10),
		sum: make([]int, 100000+10),
		cnt: make([]int, 100000+10),
	}
}

func (this *MKAverage) AddElement(num int) {
	this.ele[this.rear] = num
	this.rear++
	update(num, 100000, num, this.sum)
	update(num, 100000, 1, this.cnt)
	if this.rear-this.front > this.m {
		update(this.ele[this.front], 100000, -this.ele[this.front], this.sum)
		update(this.ele[this.front], 100000, -1, this.cnt)
		this.front++
	}
}

func (this *MKAverage) CalculateMKAverage() int {
	if this.rear-this.front < this.m {
		return -1
	}
	e1, l1 := findK(this.k, this.cnt)
	e2, l2 := findK(this.m-this.k, this.cnt)
	return ((getSum(e2, this.sum) - l2*e2) - (getSum(e1, this.sum) - l1*e1)) / this.div
}

func main() {
	fmt.Println(3 + 4>>1)
}
