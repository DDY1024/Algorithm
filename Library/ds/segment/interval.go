package main

const (
	maxn = 100010
)

var lch = func(idx int) int { return idx << 1 }
var rch = func(idx int) int { return (idx << 1) | 1 }

func build(l, r, idx int, arr []int, add []int, sum []int) {
	if l == r {
		sum[idx] = arr[l]
		add[idx] = 0
		return
	}

	mid := (l + r) >> 1
	build(l, mid, lch(idx), arr, add, sum)
	build(mid+1, r, rch(idx), arr, add, sum)
	pushUp(idx, sum) // 初始化只需要 push_up，因为不存在任何增量
}

func pushUp(idx int, sum []int) {
	sum[idx] = sum[lch(idx)] + sum[rch(idx)]
}

// len: 线段树节点代表的区间长度
func pushDown(len, idx int, add []int, sum []int) {
	if add[idx] > 0 {
		add[lch(idx)] += add[idx]
		add[rch(idx)] += add[idx]
		sum[lch(idx)] += add[idx] * (len - len>>1) // [l,(l+r)>>1]
		sum[rch(idx)] += add[idx] * (len >> 1)     // [(l+r)>>1+1,r]
		add[idx] = 0
	}
}

// L,R：待查询区间
// l,r: 线段树节点代表的区间
func query(L, R, l, r, idx int, add []int, sum []int) int {
	if L <= l && r <= R {
		return sum[idx]
	}

	// 查询 [l,r] 区间一部分，需要下推缓存的增量
	pushDown(r-l+1, idx, add, sum)

	ret, mid := 0, (l+r)>>1
	if L <= mid {
		ret += query(L, R, l, mid, lch(idx), add, sum)
	}

	if R > mid {
		ret += query(L, R, mid+1, r, rch(idx), add, sum)
	}

	// 2. 更新时已经针对该节点区间 push_up，查询时不需要再次 push_up

	return ret
}

func update(L, R, l, r, delta, idx int, add []int, sum []int) {
	if L <= l && r <= R {
		add[idx] += delta
		sum[idx] += delta * (r - l + 1)
		return
	}

	pushDown(r-l+1, idx, add, sum)

	mid := (l + r) >> 1
	if L <= mid {
		update(L, R, l, mid, delta, lch(idx), add, sum)
	}

	if R > mid {
		update(L, R, mid+1, r, delta, rch(idx), add, sum)
	}

	// 2. push_up：子节点区间的 sum 已经发生了改变，需要同时更新当前节点的 sum
	pushUp(idx, sum)
}
