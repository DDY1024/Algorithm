package main

// 参考自：https://oi-wiki.org/basic/binary/
// 应用场景
//    1. 求解凸函数（单峰函数）极值点，函数可以是连续的，也可以是离散的
// 算法优化
//    1. 三分法每次会舍弃掉左边或右边的一个区间；在选择划分点时，尽量选择区间中点
// 注意事项
//    1. 连续函数和离散函数在区间选点上的差异

const (
	eps = 1e-6
)

// 情况一：连续函数极小值
func f1(x float64) float64 {
	return x
}

func searchOne(l, r float64) float64 {
	var lmid, mid, rmid float64
	for r-l > eps {
		mid = (l + r) / 2.0
		lmid = mid - eps
		rmid = mid + eps
		if f1(lmid) < f1(rmid) {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}

// 情况二： 连续函数（极大值）
func searchTwo(l, r float64) float64 {
	var lmid, mid, rmid float64
	for r-l > eps {
		mid = (r + l) / 2.0
		lmid = mid - eps
		rmid = mid + eps
		if f1(lmid) < f1(rmid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

// 情况三：离散函数极大值
func f2(x int) int {
	return x
}

func searchThree(l, r int) int {
	var p1, p2 int
	for l < r {
		p1 = l + (r-l)/2   // 1. 区间中点
		p2 = p1 + (r-p1)/2 // 2. 区间中点 和 区间右端点的中点
		if f2(p1) < f2(p2) {
			l = p1
		} else {
			r = p2
		}
	}
	return l
}

// 情况四：离散函数极小值
func searchFour(l, r int) int {
	var p1, p2 int
	for l < r {
		p1 = l + (r-l)/2
		p2 = p1 + (r-p1)/2
		if f2(p1) < f2(p2) {
			r = p2
		} else {
			l = p1
		}
	}
	return l
}
