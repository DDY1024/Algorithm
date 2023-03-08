package main

// 参考 https://zhuanlan.zhihu.com/p/106118909

const (
	// 线段树空间占用一般为 4*n，其中 n 为表示区间的大小
	// 通常我们取值为 n<<2 + 10 即可
	maxn = 10010
)

var lch = func(idx int) int { return idx << 1 }
var rch = func(idx int) int { return (idx << 1) | 1 }

func pushUp(idx int, sum []int) {
	sum[idx] = sum[lch(idx)] + sum[rch(idx)]
}

func buildTree(l, r, idx int, arr []int, sum []int) {
	if l == r { // 叶子节点，区间大小为 1
		sum[idx] = arr[l]
		return
	}

	mid := (l + r) >> 1
	buildTree(l, mid, lch(idx), arr, sum)
	buildTree(mid+1, r, rch(idx), arr, sum)
	pushUp(idx, sum)
}

// [l, r]: 查询区间
// [s, t]: 线段树节点表示的区间
// idx: 线段树节点的索引编号
// 例如: query(l, r, 1, n, 1, sum)
func query(l, r, s, t, idx int, sum []int) int {
	if l <= s && t <= r { // 查询区间包含了线段树的节点区间，直接返回区间统计结果
		return sum[idx]
	}

	mid, ret := (s+t)>>1, 0
	if l <= mid {
		ret += query(l, r, s, mid, lch(idx), sum)
	}
	if r > mid {
		ret += query(l, r, mid+1, t, rch(idx), sum)
	}
	return ret
}

func update(pos, delta, s, t, idx int, sum []int) {
	if s == t {
		sum[idx] += delta
		return
	}

	mid := (s + t) >> 1
	if pos <= mid {
		update(pos, delta, s, mid, lch(idx), sum)
	} else {
		update(pos, delta, mid+1, t, rch(idx), sum)
	}
	pushUp(idx, sum) // 需要 push_up 进行更新
}
