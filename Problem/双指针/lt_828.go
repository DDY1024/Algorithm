package main

import "fmt"

// https://books.halfrost.com/leetcode/ChapterFour/0828.COPYRIGHT-PROBLEM-XXX/
// 枚举计算每个位置 i 处的字符对最终结果的贡献
func uniqueLetterString(S string) int {
	n, mod := len(S), 1000000000 + 7
	pos := make(map[byte][]int)
	imap := make(map[byte]int)
	for i := 0; i < n; i++ {
		pos[S[i]] = append(pos[S[i]], i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		c, idx := S[i], imap[S[i]]
		for pos[c][idx] < i {
			idx++
		}
		imap[c] = idx
		l, r := 0, 0
		if idx - 1 >= 0 {
			l = pos[c][idx-1]
		} else {
			l = -1
		}
		if idx + 1 < len(pos[c]) {
			r = pos[c][idx+1]
		} else {
			r = n
		}
		ans = (ans + (i - l) *  (r - i)) % mod // i, i+1, ..., r-1 <---> l + 1, l+2, ..., i
	}
	return ans
}

func uniqueLetterStringTwo(S string) int {
	res, left, right := 0, 0, 0
	for i := 0; i < len(S); i++ {
		left = i - 1
		for left >= 0 && S[left] != S[i] {
			left--
		}
		right = i + 1
		for right < len(S) && S[right] != S[i] {
			right++
		}
		res += (i - left) * (right - i)
	}
	return res % 1000000007
}

func main() {
	fmt.Println(uniqueLetterString("LETTER"))
	fmt.Println(uniqueLetterStringTwo("LETTER"))
	fmt.Println(uniqueLetterStringTwo("AAA"))
	fmt.Println(uniqueLetterString("AAA"))
	fmt.Println(uniqueLetterStringTwo("ABC"))
	fmt.Println(uniqueLetterString("ABC"))  // A, B, C, AB, BC, ABC = 1 + 1 + 1 + 2 + 2 + 3 = 10
}