package main

// 外部排序参考 http://data.biancheng.net/view/77.html

func quickSort(l, r int, arr []int) {
	if l >= r {
		return
	}

	idx := partition(l, r, arr)
	quickSort(l, idx-1, arr)
	quickSort(idx+1, r, arr)
}

// 划分方式
//  1. 选取 arr[l] 作为划分元素
//  2. 先寻找 arr[j] < base，赋值给 arr[i]
//  3. 后寻找 arr[i] > base，赋值给 arr[j]
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
