package main

import "sort"

// https://leetcode-cn.com/problems/subsets-ii/solution/zi-ji-ii-by-leetcode-solution-7inq/
// Tips: 重复数字部分，例如存在 3 个 2，我们选择 0，1，2，3 个 2 时，只需要考虑前面连续的选择，对于中间不选择的情况，我们直接跳过，避免重复选择。
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans, n := make([][]int, 0), len(nums)
	var dfs func(choose bool, pos, cnt int)
	arr := make([]int, n)
	dfs = func(choose bool, pos, cnt int) {
		if pos >= n {
			tmp := make([]int, cnt)
			for i := 0; i < cnt; i++ {
				tmp[i] = arr[i]
			}
			ans = append(ans, tmp)
			return
		}
		dfs(false, pos+1, cnt)
		if !choose && pos > 0 && nums[pos] == nums[pos-1] { // 前面的已经搜索过了，后面的不需要再次搜索了
			return // 注意这里是直接 return
		}
		arr[cnt] = nums[pos]
		dfs(true, pos+1, cnt+1)
	}
	dfs(false, 0, 0)
	return ans
}
