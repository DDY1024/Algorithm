package main

import (
	"math"
)

// 极角排序: https://www.cnblogs.com/aiguona/p/7248311.html
//
// math.Atan2(y, x)：值域 [-180, 180]
// 由小到大为: 第三象限 -> 第四象限 -> 第一象限 -> 第二象限

const (
	eps = 1e-4
)

type Point struct {
	x float64
	y float64
}

func less(a, b Point) bool {
	x1, x2 := math.Atan2(a.y, a.x), math.Atan2(b.y, b.x)
	if math.Abs(x1-x2) < eps {
		return a.x < b.x
	}
	return x1 < x2
}
