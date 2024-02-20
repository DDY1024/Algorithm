package main

// 题目链接：https://leetcode.cn/problems/longest-well-performing-interval/description/
// 解题报告：https://leetcode.cn/problems/longest-well-performing-interval/solutions/2110211/liang-chong-zuo-fa-liang-zhang-tu-miao-d-hysl/

// 解题思路
// 1. 问题转化：题目要求最长的连续子序列使得劳累天数严格大于不劳累天数，我们不妨将劳累天设置为 1，不劳累天设置为 -1
//		这样我们最终便转化成求解最长连续子序列且其和 "> 0"
// 2. 区间和 --> 前缀和
// 3. 利用最优方案的单调性
// 		假设前缀和 preSum[i] 和 preSum[j]，如果 preSum[i] <= preSum[j]，则以 i 为左端点的解肯定是优于 j 的，因此 j 不需要考虑
// 	遍历一遍数组，我们便得到一个严格单调递减的序列；最优答案必然以单调递减序列中的某一个点作为左端点
// 4. 正序计算 vs 逆序计算
// 		a. 正序计算，遍历到每个 i 时，需要通过二分查找单调递减序列中最小的 preSum[k]，使得 preSum[i] - preSum[k] > 0，复杂度为 O(nlogn)
//      b. 逆序计算，遍历到每个 i 时，当前队列中所有 preSum[i] - preSum[k] > 0 的 k 均为最优方案，后续的 i 不会使得其更优，
//			因此可以直接出栈，均摊下来最终复杂度为 O(n)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestWPI(hours []int) int {
	n := len(hours)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if hours[i-1] > 8 {
			preSum[i] = 1
		} else {
			preSum[i] -= 1
		}
		preSum[i] += preSum[i-1]
	}

	stk := make([]int, 0)
	stk = append(stk, 0) // 初始 0 加入队列，使得便于计算最大长度为 n 的情况
	for i := 1; i <= n; i++ {
		if preSum[i] < preSum[stk[len(stk)-1]] {
			stk = append(stk, i)
		}
	}

	ans := 0
	for i := n; i > 0; i-- {
		for len(stk) > 0 && preSum[stk[len(stk)-1]] < preSum[i] {
			ans = maxInt(ans, i-stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}
	}
	return ans
}

// 1. 由于本题前缀和只存在 -1 和 +1 操作，因此对于单调递减队列，必然是 0, -1, -2, ... 这样的形式
// 2. 单调队列中更小的 preSum，其对应的下标必然比 preSum - 1 更靠后
// 3. 由 1 和 2，我们可以得出对于某一个前缀和 preSum[i]，其最优的连续子序列必然为 preSum[i]-1 为左端点
// 4. 利用哈希存储 preSum，遍历一遍便可求解

func longestWPI2(hours []int) int {
	n := len(hours)
	mark := make(map[int]int)
	sum, ans := 0, 0
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			sum++
		} else {
			sum--
		}
		if sum > 0 { // sum{0 ~ i} 大于 > 0，整个前缀序列满足条件
			ans = maxInt(ans, i+1)
		} else {
			if p, ok := mark[sum-1]; ok {
				ans = maxInt(ans, i-p)
			}
		}
		if _, ok := mark[sum]; !ok {
			mark[sum] = i
		}
	}
	return ans
}
