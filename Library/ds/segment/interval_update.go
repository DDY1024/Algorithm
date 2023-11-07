package main

const (
	// 4 * n - 5
	// 4 * n + 10
	maxn = 100010
)

var lc = func(i int) int { return i << 1 }
var rc = func(i int) int { return (i << 1) | 1 }

// add: lazy delta
func build(l, r, i int, arr []int, add []int, sum []int) {
	if l == r {
		sum[i] = arr[l]
		add[i] = 0
		return
	}

	mid := (l + r) >> 1
	build(l, mid, lc(i), arr, add, sum)
	build(mid+1, r, rc(i), arr, add, sum)
	pushUp(i, sum)
}

func pushUp(i int, sum []int) {
	sum[i] = sum[lc(i)] + sum[rc(i)]
}

func pushDown(slen, i int, add []int, sum []int) {
	if add[i] > 0 {
		add[lc(i)] += add[i]
		add[rc(i)] += add[i]
		sum[lc(i)] += add[i] * (slen - slen>>1) // [l, (l+r)/2]
		sum[rc(i)] += add[i] * (slen >> 1)      // [(l+r)/2+1,r]
		add[i] = 0
	}
}

func query(ql, qr, l, r, i int, add []int, sum []int) int {
	if ql <= l && r <= qr {
		return sum[i]
	}

	// 1. 区间分裂
	pushDown(r-l+1, i, add, sum)

	ret, mid := 0, (l+r)>>1
	if ql <= mid {
		ret += query(ql, qr, l, mid, lc(i), add, sum)
	}

	if qr > mid {
		ret += query(ql, qr, mid+1, r, rc(i), add, sum)
	}

	// 2. 更新时该节点已经 push_up 过一次，查询时不需要再次 push_up
	// 		统计结果修正只会影响孩子节点，不会影响祖先节点
	return ret
}

func update(ql, qr, l, r, delta, i int, add []int, sum []int) {
	if ql <= l && r <= qr {
		add[i] += delta
		sum[i] += delta * (r - l + 1)
		return
	}

	pushDown(r-l+1, i, add, sum)
	mid := (l + r) >> 1
	if ql <= mid {
		update(ql, qr, l, mid, delta, lc(i), add, sum)
	}
	if qr > mid {
		update(ql, qr, mid+1, r, delta, rc(i), add, sum)
	}
	pushUp(i, sum)
}
