package shopee

// 题目链接: https://leetcode-cn.com/problems/split-array-largest-sum/
// 题目大意
// 求解最大值最小 --> 二分 + 判定

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func splitArray(nums []int, m int) int {
	n, sum, maxNum := len(nums), 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		maxNum = maxInt(maxNum, nums[i])
	}

	var check = func(limit int) bool {
		cur, cnt := 0, 0
		for i := 0; i < n; i++ {
			if cur+nums[i] > limit {
				cnt++
				cur = nums[i]
				continue
			}
			cur += nums[i]
		}
		if cur > 0 {
			cnt++
		}
		return cnt <= m
	}

	l, r, ret := maxNum, sum, sum
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			ret = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ret
}
