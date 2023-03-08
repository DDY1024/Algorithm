package main

import (
	"sort"
)

// 参考自：https://zhuanlan.zhihu.com/p/469908450

// 一、最少区间完全覆盖
//
//	给定一个长度为 m 的区间，再给出 n 条线段 [l, r]，求解使用最少的线段完全覆盖整个区间
//
// 1. 将每个区间按照左端点从小到大排序所有线段
// 2. 选择区间左端点能够覆盖【当前已经覆盖的区间右端点】中的右端点最大的线段
// 3. 按照 2 的策略不断选择相应线段，直到最终完全覆盖整个区间
//
// 相关题目
// https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/solutions/2123855/yi-zhang-tu-miao-dong-pythonjavacgo-by-e-wqry/
// 1. 灌溉花园的最少水龙头数目：https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/
// 2. 跳跃游戏：https://leetcode.cn/problems/jump-game/，求解是否可达
// 3. 跳跃游戏2：https://leetcode.cn/problems/jump-game-ii/，求解最小跳跃次数
// 4. 视频拼接：https://leetcode.cn/problems/video-stitching/
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

// 1. 排序区间，然后从头到尾依次选择最优区间，复杂度为 O(nlogn)
func minTaps(n int, ranges []int) int {
	type seg struct{ l, r int }

	segs := make([]seg, 0, n)
	for i := 0; i <= n; i++ {
		segs = append(segs, seg{maxInt(0, i-ranges[i]), minInt(n, i+ranges[i])})
	}

	sort.Slice(segs, func(i, j int) bool {
		if segs[i].l == segs[j].l {
			return segs[i].r > segs[j].r
		}
		return segs[i].l < segs[j].l
	})

	pos, cnt, idx := segs[0].r, 1, 1
	for idx <= n {
		nextPos := pos
		for idx <= n && segs[idx].l <= pos {
			nextPos = maxInt(nextPos, segs[idx].r)
			idx++
		}
		if nextPos <= pos {
			break
		}
		pos = nextPos
		cnt++
	}

	if pos >= n {
		return cnt
	}
	return -1
}

// 由于本题覆盖区间为 [0,n]，存在特殊性，存在 O(n) 解法
// 画画图就比较好理解了
func minTaps2(n int, ranges []int) int {
	rbound := make([]int, n+1) // 维护从 i 点可以跳到最右面的点坐标
	for i := 0; i <= n; i++ {
		l := maxInt(0, i-ranges[i])
		rbound[l] = maxInt(rbound[l], i+ranges[i])
	}

	cnt, curR, nextR := 0, 0, 0
	for i := 0; i < n; i++ { // 不需要遍历 n，因为 0 ~ n-1 已经能够证明 n 是否可达
		nextR = maxInt(nextR, rbound[i])
		if i == curR {
			if nextR <= i {
				return -1 // 不存在覆盖到 i+1 的方案
			}
			curR = nextR
			cnt++
		}
	}
	return cnt
}

// leetcode 1024：视频拼接
func videoStitching(clips [][]int, time int) int {
	n := len(clips)
	rBound := make([]int, time+1)
	for i := 0; i < n; i++ {
		start, end := clips[i][0], clips[i][1]
		start = minInt(start, time) // 注意越界
		end = minInt(end, time)
		rBound[start] = maxInt(rBound[start], end)
	}

	curTime, nextTime, cnt := 0, 0, 0
	for i := 0; i < time; i++ {
		nextTime = maxInt(nextTime, rBound[i])
		if i == curTime {
			if nextTime <= i {
				return -1
			}
			curTime = nextTime
			cnt++
		}
	}
	return cnt
}

// 二、最大不相交覆盖
// 		给定一个长度为m的区间，再给出n条线段的起点和终点（开区间和闭区间处理的方法是不同，这里以开区间为例），
// 	问题是从中选取尽量多的线段，使得每个线段都是独立的
//
//  1. 根据线段的【右端点进行升序排序】；右端点相同时，按照【左端点降序排序】
//  2. 最优方案必然选择第一条线段
//  3. 后续的选择为与当前线段不相交的【最小下标】线段
//  4. 上述选择的线段构成最大不相交线段集合

// 三、最小点覆盖问题
//	数轴上面有n个闭区间[a,b],取尽量少的点，使得每个区间内都至少有一个点（不同区间内含的点可以是同一个）
//
// 	1. 按照区间右端点升序排序，右端点相同时按照左端点降序排序
//  2. 最优方案必然选择第一个区间的右端点
//  3. 寻找剩下区间中第一个满足左端点 > 【当前右端点】的区间，并将其右端点设置为【当前右端点】
//  4. 依此类推，最终选择出所有的点，且满足点数量最少
