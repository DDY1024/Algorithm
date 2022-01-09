package main

// 题目链接: https://leetcode-cn.com/problems/product-of-array-except-self/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	ret[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		ret[i] = ret[i+1] * nums[i]
	}

	pNum := 1
	for i := 0; i < n; i++ {
		ret[i] = pNum
		if i+1 < n {
			ret[i] *= ret[i+1]
		}
		pNum *= nums[i]
	}
	return ret
}
