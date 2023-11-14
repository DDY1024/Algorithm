package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 此题也算是双指针解法的经典应用， a1, a2, ..., ak, b1, b2, ..., bm 两个有序序列合并后的最大长度. 枚举第一个序列的长度，然后判断第二个序列在哪儿追加。
// 由于存在单调性，所以可以采用双指针算法进行优化。
func findLengthOfShortestSubarray(arr []int) int {
	n, p1, p2 := len(arr), -1, -1
	if n < 2 {
		return 0
	}

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			p1 = i
			break
		}
	}
	for i := n-1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			p2 = i
			break
		}
	}
	if p1 == -1 {
		return 0
	}
	// 1, ..., p1 <--> p2, ..., n-1
	// two pointers
	p11, p22, maxL := p1, n-1, n-p2  // only need right result
	for p11 >= 0 {
		for p22 >= p2 && arr[p22] >= arr[p11] {
			p22--
		}
		if p22 + 1 <= n-1 {
			maxL = maxInt(maxL, p11+1+(n-1-p22))
		} else {
			maxL = maxInt(maxL, p11+1)
		}
		p11--
	}
	return n - maxL
}

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1,2,3,10,4,2,3,5}))
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3}))
	fmt.Println(findLengthOfShortestSubarray([]int{1}))
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 5, 4, 5}))
	fmt.Println(findLengthOfShortestSubarray([]int{16,10,0,3,22,1,14,7,1,12,15}))
}

// [16,10,0,3,22,1,14,7,1,12,15]
//