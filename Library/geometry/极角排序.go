package main

import (
	"fmt"
	"math"
	"sort"
)

// 关于极角排序参考：https://www.cnblogs.com/aiguona/p/7248311.html
// 关于极角排序的函数：math.Atan2

const eps = 1e-4

type Point struct {
	x float64
	y float64
}

// math.Atan2(y, x) 取值范围 [-180, 180]
func cmp(a, b Point) bool {
	x1, x2 := math.Atan2(a.y, a.x), math.Atan2(b.y, b.x)
	if math.Abs(x1-x2) < eps {
		return a.x < b.x
	}
	return x1 < x2
}

// 小 --> 大
// 第三象限、第四象限、第一象限、第二象限

func main() {
	var plist []Point
	plist = append(plist, Point{10, 0})
	plist = append(plist, Point{0, 10})
	plist = append(plist, Point{-10, 0})
	plist = append(plist, Point{0, -10})
	plist = append(plist, Point{-10, -0.1})
	sort.Slice(plist, func(i, j int) bool {
		return cmp(plist[i], plist[j])
	})
	for i := 0; i < len(plist); i++ {
		fmt.Println(plist[i])
	}
	fmt.Println(math.Atan2(10, 0))
	fmt.Println(math.Atan2(0, 10))
	fmt.Println(math.Atan2(-10, 0))
	fmt.Println(math.Atan2(0, -10))
}
