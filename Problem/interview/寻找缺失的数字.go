package main

import (
	"sort"
	"strconv"
	"strings"
)

// O(n) 时间复杂度，O(1) 空间复杂度解决
//
// 其实就是通过不断地将数字换到其对应位置即可

func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] < n && nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	// 0 ~ n-1 存在一个缺失的
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	// 0 ~ n-1 均不缺失，则 n 缺失
	return n
}

func largestNumber(nums []int) string {
	var base = func(x int) int {
		if x == 0 {
			return 10
		}
		ret := 1
		for ret <= x {
			ret *= 10
		}
		return ret
	}

	sort.Slice(nums, func(i, j int) bool { // 考察排序自定义比较规则的事情
		return nums[i]*base(nums[j])+nums[j] >= nums[j]*base(nums[i])+nums[i]
	})
	// fmt.Println(nums)

	var ret strings.Builder
	for i := 0; i < len(nums); i++ {
		ret.WriteString(strconv.Itoa(nums[i]))
	}

	// 去除前导 0
	bret := []byte(ret.String())
	for len(bret) > 1 && bret[0] == '0' {
		bret = bret[1:]
	}
	return string(bret)
}
