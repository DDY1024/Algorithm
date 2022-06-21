package main

const maxn = 100010

var lchild = func(idx int) int { return idx << 1 }
var rchild = func(idx int) int { return (idx << 1) | 1 }

func build(l, r, idx int, arr []int, add []int, sum []int) {
	if l == r {
		sum[idx] = arr[l]
		add[idx] = 0
		return
	}

	mid := (l + r) >> 1
	build(l, mid, lchild(idx), arr, add, sum)
	build(mid+1, r, rchild(idx), arr, add, sum)
	pushUp(idx, sum)
}

func pushUp(idx int, sum []int) {
	sum[idx] = sum[lchild(idx)] + sum[rchild(idx)]
}

// sl: 区间长度
func pushDown(sl, idx int, add []int, sum []int) {
	if add[idx] > 0 { // 子区间下沉操作
		add[lchild(idx)] += add[idx]
		add[rchild(idx)] += add[idx]
		sum[lchild(idx)] += add[idx] * (sl - sl>>1) // [l,(l+r)>>1]
		sum[rchild(idx)] += add[idx] * (sl >> 1)    // [(l+r)>>1+1,r]
		add[idx] = 0
	}
}

// 查询区间参数: L、R
func query(L, R, l, r, idx int, add []int, sum []int) int {
	if L <= l && r <= R {
		return sum[idx]
	}

	pushDown(r-l+1, idx, add, sum)

	ret, mid := 0, (l+r)>>1
	if L <= mid {
		ret += query(L, R, l, mid, lchild(idx), add, sum)
	}

	if R > mid {
		ret += query(L, R, mid+1, r, rchild(idx), add, sum)
	}

	// update 时已经做了 push_up，此处没有必要再做 push_up

	return ret
}

// 更新区间：L、R
func update(L, R, l, r, delta, idx int, add []int, sum []int) {
	if L <= l && r <= R {
		add[idx] += delta
		sum[idx] += delta * (r - l + 1) // 此处已经算明白了
		return
	}

	pushDown(r-l+1, idx, add, sum)

	mid := (l + r) >> 1
	if L <= mid {
		update(L, R, l, mid, delta, lchild(idx), add, sum)
	}

	if R > mid {
		update(L, R, mid+1, r, delta, rchild(idx), add, sum)
	}

	pushUp(idx, sum)
}
