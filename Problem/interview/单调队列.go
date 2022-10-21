package main

// 单调队列（双端队列）、单调栈
// 根据题目要求判定是 单调递增 or 单调递减

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	q := make([]int, 0, k)
	ret := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// 维护大小为 k 的窗口
		for len(q) > 0 && q[0] < i-k+1 {
			q = q[1:]
		}

		if i >= k-1 {
			ret = append(ret, nums[q[0]]) // 队首存储窗口内最大的元素
		}
	}

	return ret
}
