#### 一、题目链接
[最长公共子路径](https://leetcode-cn.com/problems/longest-common-subpath/)

#### 二、题目大意
求解多字符串最长公共子数组的长度。

#### 三、解题思路
本题简单形式为 [最长重复子数组](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/)，只是本题变成了多个字符串，但是多个字符串总长度 <= 10^5，因此我们仍然可以采用 `二分 + 滚动哈希` 思路来进行求解。

- `生日悖论 && 哈希冲突`：`对于一般情况，如果我们在 N 个数中可重复地随机选择 K 个数，那么在 K = sqrt(N) 的前提下，我们选择的 K 个数中存在重复的概率非常高`。
- *注意*：我们通过扩大哈希映射空间只能尽可能减少哈希冲突的可能性，这是一种取巧做法，实际上我们应该针对哈希函数相等的情况进行冲突检测。

#### 四、复杂度分析

#### 五、代码
```go
// TODO: 还是寻求正规解法吧？hash 冲突解决起来很难受
func mulMod(a, b, c int) int {
	ret := 0
	for b > 0 {
		if b&1 > 0 {
			ret = (ret + a) % c
		}
		a = a * 2 % c
		b >>= 1
	}
	return ret
}

// 此处 powMod 可能存在整数溢出
func powMod(a, b, c int) int {
	ret := 1
	for b > 0 {
		if b&1 > 0 {
			ret = mulMod(ret, a, c)
		}
		a = mulMod(a, a, c)
		b >>= 1
	}
	return ret
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Pair struct {
	h1, h2 int
}

func longestCommonSubpath(n int, paths [][]int) int {

	var (
		base1 = 100007
		base2 = 100013
		mod   = int(1e9 + 7)
		m     = len(paths)
	)

	minLen := len(paths[0])
	for i := 1; i < m; i++ {
		minLen = minInt(minLen, len(paths[i]))
	}

	var check = func(x int) bool {
		totalStats := make(map[Pair]int)
		// 双模 hash
		tmp1 := powMod(base1, x-1, mod)
		tmp2 := powMod(base2, x-1, mod)
		for i := 0; i < m; i++ {
			h1, h2 := 0, 0
			stats := make(map[Pair]bool)
			for j := 0; j < x; j++ {
				h1 = (h1*base1 + paths[i][j]) % mod
				h2 = (h2*base2 + paths[i][j]) % mod
			}
			if !stats[Pair{h1, h2}] {
				stats[Pair{h1, h2}] = true
				totalStats[Pair{h1, h2}]++
			}
			for j := x; j < len(paths[i]); j++ {
				h1 = ((h1-mulMod(paths[i][j-x], tmp1, mod)+mod)%mod*base1 + paths[i][j]) % mod
				h2 = ((h2-mulMod(paths[i][j-x], tmp2, mod)+mod)%mod*base2 + paths[i][j]) % mod
				if !stats[Pair{h1, h2}] {
					stats[Pair{h1, h2}] = true
					totalStats[Pair{h1, h2}]++
				}
			}
		}
		for _, v := range totalStats {
			if v >= m {
				return true
			}
		}
		return false
	}

	l, r, ans := 1, minLen, 0
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
```

#### 六、后缀数组
- 整理后缀数组模板，提供后缀数组解题思路。