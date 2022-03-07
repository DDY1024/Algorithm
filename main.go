package main

import "fmt"

func main() {
	fmt.Println(-2 % 9)
	fmt.Println(-1 % 9)
	fmt.Println(-10 % 9)
}

// max - min
func subArrayRanges(nums []int) int64 {
	n := len(nums)

}

// 重复元素如何处理？？？
// a, a, a
//
//
// 采用一定的方法逻辑上定义出一个最小值，确保来避免重复计算
//
// i < j && nums[i] == nums[j]
// 将下标纳入大小比较维度 --> 确保子数组的最大值/最小值唯一性
//
//
//
//
//
