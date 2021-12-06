package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(isPrime(1337))
	// largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2)
	nums := []int{3, 4, 5, 5}
	fmt.Println(sort.Search(len(nums), func(i int) bool {
		return nums[i]+2 >= 7
	}))
}

func truncateSentence(s string, k int) string {
	var buff bytes.Buffer
	buff.Grow(len(s))

	var isLetter = func(ch byte) bool {
		return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
	}

	spaceCnt := 0
	for i := 0; i < len(s) && spaceCnt < k; i++ {
		if isLetter(s[i]) {
			buff.WriteByte(s[i])
		} else {
			spaceCnt++
			if spaceCnt < k {
				buff.WriteByte(' ')
			}
		}
	}
	return buff.String()
}
