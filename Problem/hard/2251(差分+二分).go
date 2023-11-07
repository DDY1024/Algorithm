package main

import "sort"

// 方法一
// 1. 利用差分数组求解【一次性】区间操作后单点的取值
// 2. 由于数据元素很大，需要离散化处理

func fullBloomFlowers(flowers [][]int, people []int) []int {
	n, m := len(flowers), len(people)

	// 1. 排序离散化处理
	mark := make(map[int]int, n+m)
	arr := make([]int, 0, n*2+m)
	for i := 0; i < n; i++ {
		arr = append(arr, flowers[i][0])
		arr = append(arr, flowers[i][1])
	}
	for i := 0; i < m; i++ {
		arr = append(arr, people[i])
	}
	sort.Ints(arr)
	mark[arr[0]] = 0
	for idx, i := 0, 1; i < len(arr); i++ {
		if arr[i] != arr[idx] {
			idx++
			arr[idx] = arr[i]
			mark[arr[idx]] = idx
		}
	}

	// 2. 差分数组统计
	stat := make([]int, n*2+m+10)
	for i := 0; i < n; i++ {
		s, e := mark[flowers[i][0]], mark[flowers[i][1]]
		stat[s]++
		stat[e+1]--
	}
	for i := 1; i < n*2+m; i++ {
		stat[i] += stat[i-1]
	}

	// 3. 直接求解
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = stat[mark[people[i]]]
	}
	return ans
}

// 方法二
// 1. 针对每个时间点 i 花开的数目 = (起始时间 <= i 的花数) - (结束时间 < i 的花数)，便可以利用二分查找来求解计算结果

func fullBloomFlowers2(flowers [][]int, people []int) []int {
	n, m := len(flowers), len(people)
	sarr := make([]int, 0, n)
	earr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sarr = append(sarr, flowers[i][0])
		earr = append(earr, flowers[i][1])
	}
	sort.Ints(sarr)
	sort.Ints(earr)

	var calc = func(x int, arr []int) int {
		return sort.Search(n, func(i int) bool {
			return arr[i] > x
		})
	}

	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = calc(people[i], sarr) - calc(people[i]-1, earr)
	}
	return ans
}
