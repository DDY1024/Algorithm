package main

// 参考 https://zhuanlan.zhihu.com/p/106118909

const (
	// 4 * n - 5
	// 4 * n + 10
	maxn = 10010
)

var lc = func(i int) int { return i << 1 }       // 左孩子：2 * i
var rc = func(i int) int { return (i << 1) | 1 } // 右孩子：2 * i + 1

func pushUp(i int, sum []int) {
	sum[i] = sum[lc(i)] + sum[rc(i)]
}

// O(n *logn) 建树复杂度
func buildTree(l, r, i int, arr []int, sum []int) {
	if l == r {
		sum[i] = arr[l]
		return
	}

	mid := (l + r) >> 1
	buildTree(l, mid, lc(i), arr, sum)   // [l, mid]
	buildTree(mid+1, r, rc(i), arr, sum) // [mid+1, r]
	pushUp(i, sum)
}

// 例如: query(ql, qr, 1, n, 1, sum)
func query(ql, qr, l, r, i int, sum []int) int {
	// 查询区间包含了线段树节点所代表的区间，直接返回
	if ql <= l && r <= qr {
		return sum[i]
	}

	ret, mid := 0, (l+r)>>1
	if ql <= mid {
		ret += query(ql, qr, l, mid, lc(i), sum)
	}
	if qr > mid {
		ret += query(ql, qr, mid+1, r, rc(i), sum)
	}
	return ret
}

func update(pos, delta, l, r, i int, sum []int) {
	if l == r {
		sum[i] += delta
		return
	}

	mid := (l + r) >> 1
	if pos <= mid {
		update(pos, delta, l, mid, lc(i), sum)
	} else {
		update(pos, delta, mid+1, r, rc(i), sum)
	}
	pushUp(i, sum)
}
