package main

// 参考资料：https://oi-wiki.org/basic/binary/
//
// 应用场景
//    1. 求解凸函数（单峰函数）极值点（极大值、极小值），连续函数、离散函数
//
// 优化
//    1. 三分法每次会舍弃掉【左边】或【右边】的一个区间；在选择划分点时，尽量选择区间中点
//
// 注意事项
//    1. 连续函数 vs 离散函数，在区间选点上的【差异】
//

const (
	eps = 1e-4
)

func f1(x float64) float64 {
	return x
}

// 情况一：连续函数极小值
func searchOne(l, r float64) float64 {
	for r-l > eps {
		p2 := (l + r) / 2.0
		p1 := p2 - eps
		p3 := p2 + eps
		if f1(p1) < f1(p3) { // 区间中点
			r = p2
		} else {
			l = p2
		}
	}
	return l
}

// 情况二： 连续函数极大值
func searchTwo(l, r float64) float64 {
	for r-l > eps {
		p2 := (r + l) / 2.0
		p1 := p2 - eps
		p3 := p2 + eps
		if f1(p1) < f1(p3) { // 区间中点
			l = p2
		} else {
			r = p2
		}
	}
	return l
}

func f2(x int) int {
	return x
}

// 情况三：离散函数极大值
func searchThree(l, r int) int {
	for l < r {
		x1 := l + (r-l)/2   // 1. 区间中点
		x2 := x1 + (r-x1)/2 // 2. 区间中点 和 区间右端点 中间
		if f2(x1) < f2(x2) {
			l = x1
		} else {
			r = x2
		}
	}
	return l
}

// 情况四：离散函数极小值
func searchFour(l, r int) int {
	for l < r {
		x1 := l + (r-l)/2
		x2 := x1 + (r-x1)/2
		if f2(x1) < f2(x2) {
			r = x2
		} else {
			l = x1
		}
	}
	return l
}
