package main

// https://leetcode.cn/problems/sum-of-subarray-minimums/
//
// 求解子数组最小值之和
// 1. 利用单调栈处理 arr[i] 能够影响到的最左边界 L[i] 和 最右边界 R[i]
// 		则 arr[i] 对最终结果的贡献为 (i-L[i]+1)*(R[i]-i+1)*arr[i]
//
// 2. 重复元素处理
// 		增加一个下标属性，在元素值相同时，比对下标大小；这样每个元素的大小能够被唯一确定；这样计算出来的 L，R 数组不会出现
//      重复统计和遗漏统计的情况

func sumSubarrayMins(arr []int) int {
	mod := int(1e9 + 7)
	n := len(arr)
	dpl, dpr := make([]int, n), make([]int, n)
	stk := make([]int, 0)

	var less = func(i, j int) bool { // 能够唯一确定数组中每个元素的位置（重复元素）
		if arr[i] < arr[j] {
			return true
		}
		if arr[i] > arr[j] {
			return false
		}
		return i < j
	}

	dpl[0] = 0
	stk = append(stk, 0)
	for i := 1; i < n; i++ {
		for len(stk) > 0 && less(i, stk[len(stk)-1]) {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			dpl[i] = 0
		} else {
			dpl[i] = stk[len(stk)-1] + 1
		}
		stk = append(stk, i)
	}

	stk = []int{n - 1}
	dpr[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		for len(stk) > 0 && less(i, stk[len(stk)-1]) {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			dpr[i] = n - 1
		} else {
			dpr[i] = stk[len(stk)-1] - 1
		}
		stk = append(stk, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = ans + arr[i]*(i-dpl[i]+1)*(dpr[i]-i+1)%mod
		ans %= mod
	}
	return ans
}
