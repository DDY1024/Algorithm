package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 4, 5, 5}
	fmt.Println(sort.Search(len(nums), func(i int) bool {
		return nums[i]+2 >= 7
	}))
	x := 7
	x &^= 1
	fmt.Println(x)
}
