package main

// 改段求点模型
var lchild = func(idx int) int { return idx << 1 }
var rchild = func(idx int) int { return (idx << 1) | 1 }

func buildOne(l, r, idx int, add []int, arr []int) {
	if l == r {
		add[idx] = arr[l]
		return
	}
	mid := (l + r) >> 1
	buildOne(l, mid, lchild(idx), add, arr)
	buildOne(mid+1, r, rchild(idx), add, arr)
	// push_up
}

func pushDown(idx int, add []int) {
	if add[idx] > 0 { // 下沉操作
		add[lchild(idx)] += add[idx]
		add[rchild(idx)] += add[idx]
		add[idx] = 0
	}
}

func queryOne(pos, l, r, idx int, add []int) int {
	if l == r {
		return add[idx]
	}
	pushDown(idx, add)
	mid := (l + r) >> 1
	if pos <= mid {
		return queryOne(pos, l, mid, lchild(idx), add)
	}
	return queryOne(pos, mid+1, r, rchild(idx), add)
}

func updateOne(L, R, l, r, delta, idx int, add []int) {
	if L <= l && r <= R {
		add[idx] += delta
		return
	}
	pushDown(idx, add)
	mid := (l + r) >> 1
	if L <= mid {
		updateOne(L, R, l, mid, delta, lchild(idx), add)
	}
	if R > mid {
		updateOne(L, R, mid+1, r, delta, rchild(idx), add)
	}
}
