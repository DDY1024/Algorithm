package main

// 题目链接: https://leetcode.cn/problems/maximize-the-minimum-powered-city/description/

// 解题思路
// 	1. 最大最小问题: 二分枚举答案 + 判定（贪心算法）
//  2. 在确认下界为 low 时，如何建造剩余的 k 个，确保最小值 >= low？
//		假设 i 处值小于 low，则在补充时选择 i+r 处是最优的，因为此处建造能够辐射的范围最大
//  3. 如何处理数据？
// 		a. 树状数组: 改段求点模型（动态数据）
//      b. 差分数组（静态数据）

func lowbit(x int) int {
	return x & (-x)
}

func add(p, c, n int, arr []int) {
	for i := p; i <= n; i += lowbit(i) {
		arr[i] += c
	}
}

func sum(p int, arr []int) int {
	ret := 0
	for i := p; i > 0; i -= lowbit(i) {
		ret += arr[i]
	}
	return ret
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 一、树状数组解法
// func maxPower(stations []int, r int, k int) int64 {
//     n := len(stations)
//     arr := make([]int, n+1)
//     for i := 0; i < n; i++ {
//         l, r := maxInt(1, i-r+1), minInt(n, i+r+1)
//         add(l, stations[i], n, arr)
//         add(r+1, -stations[i], n, arr)
//     }

//     stations = make([]int, n+1)
//     for i := 1; i <= n; i++ {
//         stations[i] = sum(i, arr)
//     }

//     low, high := stations[1], stations[1]
//     for i := 1; i <= n; i++ {
//         low = minInt(low, stations[i])
//         high = maxInt(high, stations[i])
//     }
//     high += k

//     var check = func(x, left int) bool {
//         delta := make([]int, n+1)
//         for i := 1; i <= n; i++ {
//             cc := stations[i] + sum(i, delta)
//             if cc >= x {
//                 continue
//             }

//             if cc + left < x {
//                 return false
//             }

//             add(i, x-cc, n, delta)
//             add(i+2*r+1, cc-x, n, delta)
//             left -= x-cc
//         }
//         return true
//     }

//     ret := low
//     for low <= high {
//         mid := low + (high-low)/2
//         if check(mid, k) {
//             ret = mid
//             low = mid + 1
//         } else {
//             high = mid - 1
//         }
//     }
//     return int64(ret)
// }

// 二、差分数组解法
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		l, r := maxInt(0, i-r)+1, minInt(n-1, i+r)+1
		sum[l] += stations[i]
		if r+1 <= n {
			sum[r+1] -= stations[i]
		}
	}

	low, high := sum[1], sum[1]
	for i := 1; i <= n; i++ {
		sum[i] += sum[i-1]
		low = minInt(low, sum[i])
		high = maxInt(high, sum[i])
	}
	high += k

	var check = func(x, left int) bool {
		delta := make([]int, n+1)
		ss := 0
		for i := 1; i <= n; i++ {
			ss += delta[i] // 处理差分数组右边界减去的多余值
			if sum[i]+ss >= x {
				continue
			}

			need := x - sum[i] - ss
			if need > left {
				return false
			}

			delta[i] += need
			if i+2*r+1 <= n {
				delta[i+2*r+1] -= need
			}

			left -= need
			ss += need
		}
		return true
	}

	ret := low
	for low <= high {
		mid := low + (high-low)/2
		if check(mid, k) {
			ret = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return int64(ret)
}
