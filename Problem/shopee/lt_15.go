package shopee

import "sort"

// 题目链接: https://leetcode-cn.com/problems/3sum/
// 题目大意
// 经典的 3sum 问题，由 2sum 问题转化而来
// 1. 首先对整个数组进行排序
// 2. 枚举遍历完第一个元素后，剩余两个元素便是一个双指针问题

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := 0; i < n-2; i++ { // 如何有效地避免重复的三元组
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k, s := i+1, n-1, -nums[i]
		for j < k {
			if nums[j]+nums[k] < s {
				j++
				continue
			}
			if nums[j]+nums[k] > s {
				k--
				continue
			}
			if j < k {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ret
}
