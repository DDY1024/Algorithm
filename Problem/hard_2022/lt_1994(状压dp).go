package hard2022

// 题目链接: https://leetcode-cn.com/problems/the-number-of-good-subsets/solution/hao-zi-ji-de-shu-mu-by-leetcode-solution-ky65/
// 解题思路
// 由于 1 <= nums[i] <= 30 范围内素数个数有限，因此可以考虑采用状压 dp 方式来求解最终方案数。
// 注意对于 1 的特殊处理

func powMod(a, b, c int) int {
	ret := 1
	for b > 0 {
		if b&1 > 0 {
			ret = ret * a % c
		}
		b >>= 1
		a = a * a % c
	}
	return ret
}

// 1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29
// dp[i][j]: 状态为 i 子集元素数为 j 的方案数
// 之所以引入第二维状态 j 是为了解决一维状态下重复计数问题，算法复杂度稍微有点高
func numberOfGoodSubsets(nums []int) int {
	mod, limit := int(1e9+7), 1<<10
	dp := make([][]int, limit)
	for i := 0; i < limit; i++ {
		dp[i] = make([]int, 11)
	}

	parr := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	stats := make(map[int]int, limit)

	c1 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			c1++
			continue
		}

		mask := 0
		for j := 0; j < len(parr); j++ {
			if nums[i]%parr[j] == 0 {
				cnt := 0
				for nums[i]%parr[j] == 0 {
					nums[i] /= parr[j]
					cnt++
				}
				if cnt > 1 {
					mask = -1 // invalid state
					break
				}
				mask |= 1 << uint(j)
			}
		}
		stats[mask]++
	}

	// 注意 1<<c1 溢出问题，运用快速幂取模
	mod2 := powMod(2, c1, mod)

	var countBit = func(x int) int {
		cnt := 0
		for x > 0 {
			x &= (x - 1)
			cnt++
		}
		return cnt
	}

	fac := make([]int, 11)
	fac[0] = 1
	for i := 1; i <= 10; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	dp[0][0] = 1
	for i := 1; i < limit; i++ {
		bits := countBit(i)
		for c := 1; c <= bits; c++ {
			for j := range stats {
				if j > 0 && i&j == j {
					dp[i][c] += stats[j] * dp[i&^j][c-1] % mod
					dp[i][c] %= mod
				}
			}
			dp[i][c] = dp[i][c] * powMod(c, mod-2, mod) % mod
		}
	}

	ans := 0
	for i := 1; i < limit; i++ {
		bits := countBit(i)
		for j := 1; j <= bits; j++ {
			ans += mod2 * dp[i][j] % mod
			ans %= mod
		}
	}

	return (ans + mod) % mod
}

// 为了解决第一种方法中解决重复计数引入第二维状态导致状态计算复杂度过高的问题，我们换一种思路
// 我们在选择元素时可以按照 "从大到小" 或 "从小到大" 方式，这样便可以避免重复计算。
// dp[i][mask]: 选择 2 ~ i 范围内的数字，最终构成状态 mask 的方案数
//
// 状态转移
// dp[i][mask] = dp[i-1][mask] --> i 是一个非合法的数
// dp[i][mask] = dp[i-1][mask&^st]*freq[i] + dp[i-1][mask]
//
// 边界条件 dp[1][0] = 1<<c1

func numberOfGoodSubsetsTwo(nums []int) int {
	n := len(nums)
	mod := int(1e9 + 7)
	limit := 1 << 10
	freq := make(map[int]int, 30)
	parr := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	stats := make(map[int]int, 30)

	// dp[i][mask]: 0 ~ i 范围内的元素构成 mask 状态的素因子的方案数
	// 由于递推关系，我们在选择一个数是按照从大到小的顺序进行选择的，因此可以保证生成的子集方案数是唯一的
	dp := make([][]int, 31)
	for i := 0; i <= 30; i++ {
		dp[i] = make([]int, limit)
	}

	c1 := 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			c1++
			continue
		}

		freq[nums[i]]++
		mask, tmp := 0, nums[i]
		for j := 0; j < len(parr); j++ {
			if tmp%parr[j] == 0 {
				cnt := 0
				for tmp%parr[j] == 0 {
					tmp /= parr[j]
					cnt++
				}
				if cnt > 1 {
					mask = -1 // invalid state
					break
				}
				mask |= 1 << uint(j)
			}
		}
		stats[nums[i]] = mask
	}

	dp[1][0] = powMod(2, c1, mod) // 对于 1 的选择，我们可以选择多个或者不选，所以总共存在 2^c1 种方案
	for i := 2; i <= 30; i++ {
		for j := 0; j < limit; j++ {
			dp[i][j] = dp[i-1][j]
			mask, ok := stats[i]
			if !ok { // 判断 i 是否为一个合法的数字
				continue
			}

			// 理解这样的状态转移方程
			if j&mask == mask { // 子集关系
				dp[i][j] += dp[i-1][j&^mask] * freq[i] % mod
			}
			dp[i][j] %= mod
		}
	}

	ans := 0
	// 必须存在一个大于 1 的数字
	for i := 1; i < limit; i++ {
		ans = ans + dp[30][i]
		ans %= mod
	}
	return ans
}
