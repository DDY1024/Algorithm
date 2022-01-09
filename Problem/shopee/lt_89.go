package main

// 题目链接: https://leetcode-cn.com/problems/gray-code/
// 递归生成格雷码序列

// 递归生成格雷码
func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}
	arr := grayCode(n - 1)
	ret := make([]int, 0, 2*len(arr))
	for i := 0; i < len(arr); i++ {
		ret = append(ret, arr[i]<<1)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		ret = append(ret, arr[i]<<1|1)
	}
	return ret
}
