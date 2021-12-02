package main

import (
	"math/rand"
	"time"
)

// 题目链接：https://leetcode-cn.com/problems/random-flip-matrix/
// 解题思路: 利用 hash 映射减少随机数构造的成本，确保一次随机调用即可

// 举例说明
// 利用 hash 做映射替换，比较巧妙
// 0, 1, 2, 3, 4  --> 3
// 0, 1, 2, 4 --> 1
// 0, 4, 2 --> 0
// 2, 4 --> 4
// 2 --> 2

type Solution struct {
	mark  map[int]int
	total int
	m     int
	n     int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Constructor(m int, n int) Solution {
	mark := make(map[int]int, 1000)
	return Solution{
		mark:  mark,
		total: m * n,
		m:     m,
		n:     n,
	}
}

func (this *Solution) Flip() []int {
	// 如果当前随机选择的 x 不是尾部元素，在将 x 调整成尾部元素，继续参与下一轮随机选择
	// 恰好为尾部元素，则删除就删除了，不影响后续随机选择
	x := rand.Intn(this.total) // hash 映射维护 x 被哪个数替代了，这样便可以只触发一次随机函数调用;
	this.total--

	get := func(x int) int {
		if r, ok := this.mark[x]; ok {
			return r
		}
		return x
	}

	r := get(x)
	this.mark[x] = get(this.total)
	return []int{r / this.n, r % this.n}
}

func (this *Solution) Reset() {
	this.total = this.m * this.n
	for k := range this.mark {
		delete(this.mark, k)
	}
}
