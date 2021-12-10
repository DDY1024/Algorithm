package main

import "sort"

// 熟悉扫描线算法思想解题思路
// 扫描线算法求解参考：https://leetcode-cn.com/problems/perfect-rectangle/solution/gong-shui-san-xie-chang-gui-sao-miao-xia-p4q4/
// y1: 下端点
// y2: 上端点
// d: 属于矩形左边的竖边还是右边的竖边, -1: 左边 1: 右边
type Edge struct {
	y1, y2 int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isRectangleCover(rectangles [][]int) bool {
	n := len(rectangles)
	left := make(map[int][]Edge, n)
	right := make(map[int][]Edge, n)
	markX := make(map[int]struct{}, n)
	minX, maxX, minY, maxY := 0x3f3f3f3f, -0x3f3f3f3f, 0x3f3f3f3f, -0x3f3f3f3f
	for i := 0; i < n; i++ {
		x1, y1, x2, y2 := rectangles[i][0], rectangles[i][1], rectangles[i][2], rectangles[i][3]
		if _, ok := left[x1]; !ok {
			left[x1] = make([]Edge, 0, 100)
		}
		left[x1] = append(left[x1], Edge{y1, y2})
		if _, ok := right[x2]; !ok {
			right[x2] = make([]Edge, 0, 100)
		}
		right[x2] = append(right[x2], Edge{y1, y2})
		maxX = maxInt(maxX, x2)
		minX = minInt(minX, x1)
		minY = minInt(minY, y1)
		maxY = maxInt(maxY, y2)
		markX[x1] = struct{}{}
		markX[x2] = struct{}{}
	}

	// 按照端点从低到高进行排序
	for x := range left {
		sort.Slice(left[x], func(i, j int) bool {
			return left[x][i].y2 < left[x][j].y2
		})
	}
	for x := range right {
		sort.Slice(right[x], func(i, j int) bool {
			return right[x][i].y2 < right[x][j].y2
		})
	}

	// 先处理两条边界边: 长度相等且只出现一次，避免出现重叠情况
	arr := left[minX]
	ty1, ty2 := left[minX][0].y1, left[minX][0].y1
	for i := 0; i < len(arr); i++ {
		if ty2 == arr[i].y1 { // 注意所有区间不能存在重叠的情况，因此此处用 == 判断而不是 >=
			ty2 = arr[i].y2
		} else {
			return false
		}
	}
	if ty1 != minY || ty2 != maxY {
		return false
	}

	arr = right[maxX]
	ty1, ty2 = right[maxX][0].y1, right[maxX][0].y1
	for i := 0; i < len(arr); i++ {
		if ty2 == arr[i].y1 { // 注意所有区间不能存在重叠的情况，因此此处用 == 判断而不是 >=
			ty2 = arr[i].y2
		} else {
			return false
		}
	}
	if ty1 != minY || ty2 != maxY {
		return false
	}

	delete(markX, minX)
	delete(markX, maxX)
	// 不断地合并区间，判断两个数组代表的合并区间是否相同
	// 用 == 判断衔接，同时处理区间相交和重叠情况
	for x := range markX {
		arr1 := left[x]
		arr2 := right[x]
		i, j := 0, 0
		for i < len(arr1) && j < len(arr2) {
			ty1, ty2 := arr1[i].y1, arr1[i].y1
			tz1, tz2 := arr2[j].y1, arr2[j].y1
			for i < len(arr1) && ty2 == arr1[i].y1 { // 注意所有区间不能存在重叠的情况，因此此处用 == 判断而不是 >=
				ty2 = arr1[i].y2
				i++
			}
			for j < len(arr2) && tz2 == arr2[j].y1 { // 注意所有区间不能存在重叠的情况，因此此处用 == 判断而不是 >=
				tz2 = arr2[j].y2
				j++
			}
			if !(ty1 == tz1 && ty2 == tz2) {
				return false
			}
		}
		if i < len(arr1) || j < len(arr2) {
			return false
		}
	}
	return true
}
