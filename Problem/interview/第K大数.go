package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// index 取值范围 [0, n)，元素按照从小到大顺序排列
func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	}

	if q < index {
		return quickSelect(a, q+1, r, index)
	}

	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

// 返回划分元素的下标索引
func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
