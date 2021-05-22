package main

// 参考: https://oi-wiki.org/math/cantor/
// 复杂度: O(n^2) O(nlogn)
// 优化：康托展开中我们需要求解剩下的数中比当前数小的个数，此处我们是可以用树状数组进行维护的，达到 O(nlogn) 复杂度

var N int
var factor []int

func init() {
	N = 9
	factor = make([]int, N+1)
	factor[0] = 1
	for i := 1; i <= N; i++ {
		factor[i] = factor[i-1] * i
	}
}

// 1. 康托展开: 求解 1 ~ n 排列的排名
func CantorExpansion(perm []int) int {
	n := len(perm)
	mark := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		mark[i] = true
	}

	rank := 0
	for i := 0; i < n; i++ {
		lessCnt := 0
		for j := 1; j < perm[i]; j++ {
			if mark[j] {
				lessCnt++
			}
		}
		rank += factor[lessCnt]
		mark[perm[i]] = false
	}
	return rank + 1
}

// 2. 逆康托展开: 给定排名，还原排列
func InverseCantorExpansion(n, rank int) []int {
	result := make([]int, 0, n)
	mark := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		mark[i] = true
	}

	rank -= 1
	for i := 1; i <= n; i++ {
		lessCnt, choose := rank/factor[n-i]+1, -1
		rank -= (lessCnt - 1) * factor[n-i]
		for j := 1; j <= n && lessCnt > 0; j++ {
			if mark[j] {
				lessCnt--
				choose = j
			}
		}
		result = append(result, choose)
		mark[choose] = false
	}
	return result
}
