package main

import (
	"math"
	"sort"
)

// 题目链接: https://leetcode-cn.com/problems/maximum-number-of-visible-points/
// 解题思路
// 1. 转化成极角坐标系，按照极角排序
// 2. 筛选出指定区间内点的最大数目
// 题解可以参考: https://leetcode-cn.com/problems/maximum-number-of-visible-points/solution/gong-shui-san-xie-qiu-ji-jiao-ji-he-ti-b-0bid/

// 注意: 极角取值范围是一个循环，因此我们在求解的时候需要考虑循环处理。
// 提供一种思路是将数组元素复制一份追加，整体弧度 + 2*pi

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func visiblePoints(points [][]int, angle int, location []int) int {
	sameCnt, sx, sy, n := 0, location[0], location[1], len(points)
	arcList := make([]float64, 0, 2*n)
	for i := 0; i < n; i++ {
		if sx == points[i][0] && sy == points[i][1] { // 坐标相同的点提前摘出来，避免一些额外计算
			sameCnt++
			continue
		}
		// [-180, 180]
		arcList = append(arcList, math.Atan2(float64(points[i][1]-sy), float64(points[i][0]-sx)))
	}

	arcAngle := math.Pi * float64(angle) / 180.0 // 角度转弧度

	n = len(arcList)
	for i := 0; i < n; i++ {
		// 180,-180 边界处理
		arcList = append(arcList, arcList[i]+2*math.Pi) // + 2 * pi，
	}
	n = 2 * n
	// 极角从小到大排序，筛选出固定区间内的最大点数
	sort.Slice(arcList, func(i, j int) bool {
		return arcList[i] < arcList[j]
	})

	i, j, ans := 0, 0, sameCnt
	for j < n {
		for i < j && arcList[j]-arcList[i] > arcAngle {
			i++
		}
		ans = maxInt(ans, sameCnt+j-i+1)
		j++
	}
	return ans
}

// math.Atan2(0.0, 0.0) = 0
// [0,180] [-180,0]
