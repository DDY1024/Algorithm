package main

import "fmt"

func main() {
	fmt.Println(-2 ^ 3)
}

func countVowelSubstrings(word string) int {
	n, cnt := len(word), 0

	// 11111
	yy := []byte{'a', 'e', 'i', 'o', 'u'}
	var isValid = func(s string) bool {
		mask := 0
		for i := 0; i < len(s); i++ {
			ok := false
			for j := 0; j < 5; j++ {
				if s[i] == yy[j] {
					ok = true
					mask |= 1 << uint(j)
					break
				}
			}
			if !ok {
				return false
			}
		}
		return mask == 31
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isValid(word[i : j+1]) {
				cnt++
			}
		}
	}
	return cnt
}

func countVowels(word string) int64 {
	n := len(word)
	pSum := make([]int, n)
	// ppSum := make([]int, n)

	var isyy = func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	for i := 0; i < n; i++ {
		if isyy(word[i]) {
			pSum[i]++
		}
		if i > 0 {
			pSum[i] += pSum[i-1]
		}
	}

	cur := pSum[0]
	ans := pSum[0]
	for i := 1; i < n; i++ {
		ans += (i + 1) * pSum[i]
		ans -= cur
		cur += pSum[i]
	}
	return int64(ans)
}

func minimizedMaximum(n int, quantities []int) int {
	m := len(quantities)

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var can = func(x int) bool {
		idx, need, r := 0, 0, quantities[0]
		for {
			if r == 0 {
				idx++
				if idx >= m {
					break
				}
				r = quantities[idx]
			}
			r -= minInt(x, r)
			need++
		}
		return need <= n
	}

	l, r, ret := 1, 100000, 0x3f3f3f3f
	for l <= r {
		mid := l + (r-l)/2
		if can(mid) {
			ret = minInt(ret, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if ret >= 0x3f3f3f3f {
		return -1
	}
	return ret
}

// 本题解题思路需要从数据范围中挖掘
// 1. 10 <= timej, maxTime <= 100 意味着最终结果最多经过 10 条边
// 2. 每个顶点最多有四条出边
// 3. 加上一定的剪枝策略，我们完全可以在较短时间内进行求解
type Edge struct {
	v, w int
}

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n, m := len(values), len(edges)
	adj := make([][]*Edge, n)
	for i := 0; i < m; i++ {
		u, v, w := edges[i][0], edges[i][1], edges[i][2]
		adj[u] = append(adj[u], &Edge{v, w})
		adj[v] = append(adj[v], &Edge{u, w})
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0
	path := make(map[int]int, 20)
	path[0]++
	var dfs func(u, cost int)
	dfs = func(u, cost int) {
		if u == 0 {
			tmp := 0
			for u, c := range path {
				if c > 0 {
					tmp += values[u]
				}
			}
			ans = maxInt(ans, tmp)
			return
		}

		for _, e := range adj[u] {
			if cost+e.w > maxTime {
				continue
			}
			path[e.v]++
			dfs(e.v, cost+e.w)
			path[e.v]--
			if path[e.v] == 0 {
				delete(path, e.v)
			}
		}
	}
	dfs(0, 0)
	return ans
}
