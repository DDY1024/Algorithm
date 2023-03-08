package main

// 1. 下标索引范围 [0,n-1)
// 2. 第 K 大对应的下标索引为 n-k
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := partition(a, l, r)
	if q == index {
		return a[q]
	}

	if q < index {
		return quickSelect(a, q+1, r, index)
	}

	return quickSelect(a, l, q-1, index)
}

func partition(arr []int, l, r int) int {
	i, j, base := l, r, arr[l]
	for i < j {
		// 1. 先 j
		for i < j && arr[j] >= base { // >=
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		// 2. 后 i
		for i < j && arr[i] <= base { // <=
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = base
	return i
}
