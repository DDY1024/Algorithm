package main

import (
	"fmt"
	"sort"
)

// 题目链接: https://leetcode-cn.com/problems/maximum-running-time-of-n-computers/
// 解题思路
// 1. 很容易发现答案满足单调性，即 x 成立，则 x-1 必然成立
// 2. 因此我们可以二分枚举答案，然后判断该答案是否可行(重点想好怎么进行判定当前答案是否可行)
// 另外一个大佬的题解可以参考: https://leetcode-cn.com/problems/maximum-running-time-of-n-computers/solution/liang-chong-jie-fa-er-fen-da-an-pai-xu-t-grd8/

func maxRunTime(n int, batteries []int) int64 {
	m, sum := len(batteries), 0
	sort.Ints(batteries)
	for i := 0; i < m; i++ {
		sum += batteries[i]
	}

	var check = func(x int) bool {
		left, sum := n, 0
		for i := m - 1; i >= 0; i-- {
			// 1. 对于 >= x 的电池只能被一台电脑占有
			if batteries[i] >= x {
				left--
				continue
			}
			sum += batteries[i]
		}
		// 2. 剩下的电脑 + 电池，我们 1 分钟一分钟的分，如果平均数 >= x，说明电池数量是大于电脑数量，且肯定是可以供剩下的电脑充电 x 分钟的
		if sum >= x*left {
			return true
		}
		return false
	}

	l, r, ret := 1, sum/n, -1
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return int64(ret)
}

func main() {
	fmt.Println(maxRunTime(9, []int{18, 54, 2, 53, 87, 31, 71, 4, 29, 25}))
}
