package main

func nextPermutaion(arr []int) {
	n, idx := len(arr), len(arr)-2
	for idx >= 0 && arr[idx] >= arr[idx+1] { // 排除掉重复元素的影响
		idx--
	}
	if idx < 0 { // 最大排序序列
		return
	}

	for i := n - 1; i > idx; i-- {
		if arr[i] > arr[idx] {
			arr[i], arr[idx] = arr[idx], arr[i]
			break
		}
	}

	for i, j := idx+1, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
