package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand.Seed(time.Now().UnixNano()) depcrated
}

func Shuffle(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		j := r.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func stdShuffle(arr []int) []int {
	r.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(stdShuffle(arr))
}
