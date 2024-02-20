package main

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

// https://leetcode-cn.com/problems/maximum-number-of-groups-getting-fresh-donuts/
// 1. 首先容易知道对于 > batch_size 的 group，我们只需要关注其余数即可
// 2. 对于余数为 0 的 group 我们优先招待，因为放到后面只会受到别的 group 影响，放到前面却不会影响别的 group
// 3. 对于剩下的 group，我们优先选择两两配对能够凑够 batch_size 的情况，保证其中的一个 group 是最优的
// 4. 对于剩下的更多 group 凑成 batch_size 的情况，我们采用 dfs 进行记忆化搜索。(此时最多剩下 4 个组)

func maxHappyGroups(batchSize int, groups []int) int {
	ans, n := 0, len(groups)
	mod := make([]int, batchSize)
	for i := 0; i < n; i++ {
		groups[i] %= batchSize
		if groups[i] == 0 {
			ans++
			continue
		}
		mod[groups[i]]++
	}

	for i := 1; i <= batchSize/2; i++ {
		if i*2 == batchSize {
			ans += mod[i] / 2
			mod[i] %= 2
		} else {
			cc := minInt(mod[i], mod[batchSize-i])
			ans += cc
			mod[i] -= cc
			mod[batchSize-i] -= cc
		}
	}

	// 此时最多剩余 4 组，每组上限为 30，因此我们采用 bit 来表示状态
	mask, shiftSize := 0x1f, 5
	left, stat, sum := make([]int, 0), 0, 0
	for i := 1; i < batchSize; i++ {
		if mod[i] > 0 {
			left = append(left, i)
			stat |= mod[i] << uint((len(left)-1)*shiftSize)
			sum += i * mod[i]
		}
	}

	dp := make(map[int]int)
	dp[0] = 0
	var dfs func(stat, sum int) int
	dfs = func(stat, sum int) int {
		if x, ok := dp[stat]; ok {
			return x
		}
		ans := 0
		for i := 0; i < len(left); i++ {
			cc := (stat >> uint(i*shiftSize)) & mask
			if cc > 0 {
				if (sum-left[i])%batchSize == 0 {
					ans = maxInt(ans, dfs(stat-(1<<uint(i*shiftSize)), sum-left[i])+1)
				} else {
					ans = maxInt(ans, dfs(stat-(1<<uint(i*shiftSize)), sum-left[i]))
				}
			}
		}
		dp[stat] = ans
		return ans
	}
	ans += dfs(stat, sum)
	return ans
}

// func main() {
// 	fmt.Println(maxHappyGroups(9, []int{1, 2, 6, 1, 2, 6}))
// }

// 4 种情况, 30
// 7^2 * 8^2
// 7, 7, 8, 8
//
// 7^2*8^2
// 复杂度分析是门艺术活
// 估算好复杂度选择合适的方法进行求解
// 4^
//
// dp[30][30][30][30]
// c1, c2, c3, c4 --> sum,
//
