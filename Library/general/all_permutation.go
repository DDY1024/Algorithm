package main

import "fmt"

func allPermutation(arr []int) {
	n := len(arr)

	var do func(pos int)
	do = func(pos int) {
		if pos >= n {
			fmt.Println(arr)
			return
		}

		mark := make(map[int]bool, 0)
		for i := pos; i < n; i++ {
			if mark[arr[i]] {
				continue
			}

			arr[pos], arr[i] = arr[i], arr[pos]
			do(pos + 1)
			arr[pos], arr[i] = arr[i], arr[pos]
			mark[arr[i]] = true
		}
	}
	do(0)
}
