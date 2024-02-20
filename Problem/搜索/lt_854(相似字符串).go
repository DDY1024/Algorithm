package main

// 题目链接: https://leetcode.cn/problems/k-similar-strings/
//
// 由于题目数据范围为 20，考虑采用 dfs + 剪枝策略进行求解

// 官方题解：https://leetcode.cn/problems/k-similar-strings/

// 剪枝思路:
// 针对 s 和 t 最少交换次数为 (k+1)/2，其中 k 为相同位置且字母不同的总次数；最少交换次数是两两交换，恰好满足相同

func kSimilarity(s1 string, s2 string) int {
	n := len(s1)
	bs, bt := make([]byte, 0, n), make([]byte, 0, n)

	// 只处理不同的位置即可
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			bs = append(bs, s1[i])
			bt = append(bt, s2[i])
		}
	}

	n = len(bs)
	if n == 0 {
		return 0
	}

	// 最小交换次数 (不同位置次数+1)/2
	var minSwap = func(idx int) int {
		diff := 0
		for i := idx; i < len(bs); i++ {
			if bs[i] != bt[i] {
				diff++
			}
		}
		return (diff + 1) / 2
	}

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ans := n - 1
	var dfs func(pos, cost int)
	dfs = func(pos, cost int) {
		if pos >= n {
			ans = minInt(ans, cost)
			return
		}

		if bs[pos] == bt[pos] {
			dfs(pos+1, cost)
			return
		}

		// 剪枝优化: 当前交换次数 + 剩余最小交换次数 >= ans，当前方案不可能更优，直接剪枝掉
		if cost+minSwap(pos) >= ans {
			return
		}

		for i := pos + 1; i < n; i++ {
			if bs[i] == bt[pos] {
				bs[pos], bs[i] = bs[i], bs[pos]
				dfs(pos+1, cost+1)
				bs[pos], bs[i] = bs[i], bs[pos] // 回溯
			}
		}
	}
	dfs(0, 0)
	return ans
}
