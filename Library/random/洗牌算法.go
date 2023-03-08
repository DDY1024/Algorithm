package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 每个数在每个位置出现的概率均为 1/n
func Shuffle(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}
