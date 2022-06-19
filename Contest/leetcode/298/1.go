package main

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func greatestLetter(s string) string {
	c1, c2 := make([]int, 26), make([]int, 26)
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			c1[int(s[i]-'a')]++
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			c2[int(s[i]-'A')]++
		}
	}

	for i := 25; i >= 0; i-- {
		if c1[i] > 0 && c2[i] > 0 {
			return string([]byte{byte('A' + i)})
		}
	}
	return ""
}
