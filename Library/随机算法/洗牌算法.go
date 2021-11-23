package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(n int, swap func(i, j int))
}

// 每个数在每个位置出现的概率均等为 1/n
func Shuffle(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
