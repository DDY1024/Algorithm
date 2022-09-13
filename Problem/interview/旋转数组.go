package main

import (
	"math/rand"
	"time"
)

// 旋转数组: O(1) 空间复杂度原地旋转
// 利用矩阵的置换运算
// (A^t,B^t)^t = B^t^t,A^t^t=BA

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	var reverse = func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	reverse(nums[:n-k])
	reverse(nums[n-k:])
	reverse(nums)
}

type Solution struct {
	ori  []int
	data []int
	size int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	ori := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ori[i] = nums[i]
	}
	return Solution{
		ori:  ori,
		data: nums,
		size: len(nums),
	}
}

func (this *Solution) Reset() []int {
	return this.ori
}

func (this *Solution) Shuffle() []int {
	// rand.Shuffle(this.size, func(i, j int) {
	// 	this.data[i], this.data[j] = this.data[j], this.data[i]
	// })
	for i := this.size - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		this.data[i], this.data[j] = this.data[j], this.data[i]
	}
	return this.data
}
