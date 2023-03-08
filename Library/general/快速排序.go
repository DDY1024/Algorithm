package main

// 参考自：https://blog.csdn.net/wthfeng/article/details/78037228

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	i := partition(low, high, arr)
	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}

// 选择一种最容易理解的划分方式
// 1. 选取 arr[l] 作为切分元素
// 2. 先寻找 arr[j] < base，赋值给 arr[i]
// 2. 后寻找 arr[i] > base，赋值给 arr[j]
func partition(l, r int, arr []int) int {
	i, j, base := l, r, arr[l]
	for i < j {
		for i < j && arr[j] >= base {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= base {
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
